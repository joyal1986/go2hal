package chef

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-chef/chef"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/util"
	"gopkg.in/kyokomi/emoji.v1"
	"strings"
	"time"
	"github.com/weAutomateEverything/go2hal/telegram"
)

type Service interface {
	sendDeliveryAlert(ctx context.Context, message string) error
	FindNodesFromFriendlyNames(recipe, environment string) ([]Node, error)
	SendKeyboardRecipe(ctx context.Context, message string) error
	SendKeyboardEnvironment(ctx context.Context, message string) error
	SendKeyboardNodes(ctx context.Context, recipe, environment, message string) error
}

type service struct {
	alert     alert.Service
	chefStore Store
	telegram  telegram.Service
}

func NewService(alert alert.Service, chefStore Store, telegram telegram.Service) Service {
	s := &service{alert, chefStore, telegram}
	go func() {
		s.monitorQuarentined()
	}()
	return s
}

func (s *service) sendDeliveryAlert(ctx context.Context, message string) error {
	var dat map[string]interface{}

	message = strings.Replace(message, "\n", "\\n", -1)

	if err := json2.Unmarshal([]byte(message), &dat); err != nil {
		s.alert.SendError(ctx, fmt.Errorf("delivery - error unmarshalling: %s", message))
		return err
	}

	attachments := dat["attachments"].([]interface{})

	body := dat["text"].(string)
	bodies := strings.Split(body, "\n")
	url := bodies[0]
	url = strings.Replace(url, "<", "", -1)
	url = strings.Replace(url, ">", "", -1)

	parts := strings.Split(url, "|")

	//Loop though the attachmanets, there should be only 1
	var buffer bytes.Buffer
	buffer.WriteString(emoji.Sprint(":truck:"))
	buffer.WriteString(" ")
	buffer.WriteString("*chef Delivery*\n")

	if len(bodies) > 1 {
		buildDeliveryEnent(&buffer, bodies[1])
	} else {
		buffer.WriteString(emoji.Sprintf(":rage1: New Code Review \n"))
	}

	util.Getfield(attachments, &buffer)

	buffer.WriteString("[")
	buffer.WriteString(parts[1])

	buffer.WriteString("](")
	buffer.WriteString(parts[0])
	buffer.WriteString(")")

	s.alert.SendAlert(ctx, buffer.String())
	return nil
}

func buildDeliveryEnent(buffer *bytes.Buffer, body string) {
	if strings.Contains(body, "failed") {
		buffer.WriteString(emoji.Sprint(":interrobang:"))

	} else {
		switch body {
		case "Delivered stage has completed for this change.":
			buffer.WriteString(emoji.Sprint(":+1:"))

		case "Change Delivered!":
			buffer.WriteString(emoji.Sprint(":white_check_mark:"))

		case "Acceptance Passed. Change is ready for delivery.":
			buffer.WriteString(emoji.Sprint(":ok_hand:"))

		case "Change Approved!":
			buffer.WriteString(emoji.Sprint(":white_check_mark:"))

		case "Verify Passed. Change is ready for review.":
			buffer.WriteString(emoji.Sprint(":mag_right:"))
		}
	}
	buffer.WriteString(" ")

	buffer.WriteString(body)
	buffer.WriteString("\n")
}

func (s *service) monitorQuarentined() {
	for {
		s.checkQuarentined()
		time.Sleep(30 * time.Minute)
	}
}
func (s *service) checkQuarentined() {
	recipes, err := s.chefStore.GetRecipes()
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return
	}

	env, err := s.chefStore.GetChefEnvironments()
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return
	}

	for _, r := range recipes {
		for _, e := range env {
			nodes, _ := s.FindNodesFromFriendlyNames(r.FriendlyName, e.FriendlyName)
			for _, n := range nodes {
				if strings.Index(n.Environment, "quar") > 0 {
					s.alert.SendAlert(context.TODO(), emoji.Sprintf(":hospital: *Node Quarantined* \n node %v has been placed in environment %v. Application %v ", n.Name, strings.Replace(n.Environment, "_", " ", -1), r.FriendlyName))
				}
			}
		}
	}

}

func (s *service) FindNodesFromFriendlyNames(recipe, environment string) ([]Node, error) {
	chefRecipe, err := s.chefStore.GetRecipeFromFriendlyName(recipe)
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return nil, err
	}

	chefEnv, err := s.chefStore.GetEnvironmentFromFriendlyName(environment)
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return nil, err
	}

	client, err := s.getChefClient()
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return nil, err
	}

	query, err := client.Search.NewQuery("node", fmt.Sprintf("recipe:%s AND chef_environment:%s", chefRecipe, chefEnv))
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return nil, err
	}

	part := make(map[string]interface{})
	part["name"] = []string{"name"}
	part["chef_environment"] = []string{"chef_environment"}

	res, err := query.DoPartial(client, part)
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return nil, err
	}

	result := make([]Node, res.Total)

	for i, x := range res.Rows {
		s := x.(map[string]interface{})
		data := s["data"].(map[string]interface{})
		name := data["name"].(string)
		env := data["chef_environment"].(string)
		result[i] = Node{Name: name, Environment: env}
	}

	return result, nil

}

func (s *service) getChefClient() (client *chef.Client, err error) {
	c, err := s.chefStore.GetChefClientDetails()
	if err != nil {
		return nil, err
	}
	client, err = connect(c.Name, c.Key, c.URL)
	return client, err
}

func connect(name, key, url string) (client *chef.Client, err error) {
	client, err = chef.NewClient(&chef.Config{
		Name:    name,
		Key:     key,
		BaseURL: url,
		SkipSSL: true,
	})
	return
}

type Node struct {
	Name        string
	Environment string
}

func (s *service) SendKeyboardRecipe(ctx context.Context, message string) error {
	recipes, err := s.chefStore.GetRecipes()
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return err
	}
	l := make([]string, len(recipes))
	for x, i := range recipes {
		l[x] = i.FriendlyName
	}
	alertGroup, err := s.alert.AlertGroup(ctx)
	s.telegram.SendKeyboard(ctx, l, message, alertGroup)
	return err;
}
func (s *service) SendKeyboardEnvironment(ctx context.Context, message string) error {
	environments, err := s.chefStore.GetChefEnvironments()
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return err
	}

	l := make([]string, len(environments))
	for x, i := range environments {
		l[x] = i.FriendlyName
	}
	alertGroup, err := s.alert.AlertGroup(ctx)
	s.telegram.SendKeyboard(ctx, l, message, alertGroup)
	return err;
}
func (s *service) SendKeyboardNodes(ctx context.Context, recipe, environment, message string) error {
	nodes, err := s.FindNodesFromFriendlyNames(recipe, environment)
	if err != nil {
		s.alert.SendError(context.TODO(), err)
		return err
	}

	l := make([]string, len(nodes))
	for x, i := range nodes {
		l[x] = i.Name
	}
	alertGroup, err := s.alert.AlertGroup(ctx)
	s.telegram.SendKeyboard(ctx, l, message, alertGroup)
	return err;
}
