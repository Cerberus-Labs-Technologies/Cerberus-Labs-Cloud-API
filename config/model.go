package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

type Config struct {
	FileName string
	Data     map[string]interface{} `json:"data"`
}

func (c *Config) Get(key string) interface{} {
	if val, ok := c.Data[key]; ok {
		return val
	}

	// If the key is not found at the top level of the Data map,
	// try to access it as a nested field using dot notation
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		return nil
	}
	val := c.Data[parts[0]]
	for _, part := range parts[1:] {
		nested, ok := val.(map[string]interface{})
		if !ok {
			return nil
		}
		val = nested[part]
	}
	return val
}

func (c *Config) Set(key string, value interface{}) error {
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		c.Data[key] = value
		return nil
	}
	val, ok := c.Data[parts[0]].(map[string]interface{})
	if !ok {
		val = make(map[string]interface{})
		c.Data[parts[0]] = val
	}
	for _, part := range parts[1 : len(parts)-1] {
		nested, ok := val[part].(map[string]interface{})
		if !ok {
			nested = make(map[string]interface{})
			val[part] = nested
		}
		val = nested
	}
	val[parts[len(parts)-1]] = value
	return nil
}

func (c *Config) Delete(key string) error {
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		delete(c.Data, key)
		return nil
	}
	val, ok := c.Data[parts[0]].(map[string]interface{})
	if !ok {
		return fmt.Errorf("cannot delete non-existent key: %s", key)
	}
	for _, part := range parts[1 : len(parts)-1] {
		nested, ok := val[part].(map[string]interface{})
		if !ok {
			return fmt.Errorf("cannot delete non-existent key: %s", key)
		}
		val = nested
	}
	delete(val, parts[len(parts)-1])
	return nil
}

func (c *Config) GetInt(key string) int {
	return c.Get(key).(int)
}

func (c *Config) GetFloat(key string) float64 {
	return c.Get(key).(float64)
}

func (c *Config) GetString(key string) string {
	return c.Get(key).(string)
}

func (c *Config) GetBool(key string) bool {
	return c.Get(key).(bool)
}

func (c *Config) GetMap(key string) map[string]interface{} {
	return c.Get(key).(map[string]interface{})
}

func (c *Config) GetArray(key string) []interface{} {
	return c.Get(key).([]interface{})
}

func (c *Config) IsEmpty(key string) bool {
	// Check if the key exists or is empty array
	if _, ok := c.Data[key]; !ok || len(c.Data[key].([]interface{})) == 0 {
		return true
	}
	return false
}

func (c *Config) GetArrayString(key string) []string {
	var array []string
	for _, v := range c.Get(key).([]interface{}) {
		array = append(array, v.(string))
	}
	return array
}

func (c *Config) GetArrayInt(key string) []int {
	var array []int
	for _, v := range c.Get(key).([]interface{}) {
		array = append(array, v.(int))
	}
	return array
}

func (c *Config) GetArrayFloat(key string) []float64 {
	var array []float64
	for _, v := range c.Get(key).([]interface{}) {
		array = append(array, v.(float64))
	}
	return array
}

func (c *Config) GetArrayBool(key string) []bool {
	var array []bool
	for _, v := range c.Get(key).([]interface{}) {
		array = append(array, v.(bool))
	}
	return array
}

func (c *Config) GetArrayMap(key string) []map[string]interface{} {
	var array []map[string]interface{}
	for _, v := range c.Get(key).([]interface{}) {
		array = append(array, v.(map[string]interface{}))
	}
	return array
}

func (c *Config) AddToArray(key string, value interface{}) error {
	arr, ok := c.Data[key].([]interface{})
	if !ok {
		// If the key doesn't exist or the value is not an array,
		// create a new array and set it as the value for the key.
		arr = make([]interface{}, 0)
	}
	c.Data[key] = append(arr, value)
	return nil
}

func (c *Config) RemoveFromArray(key string, value interface{}, valueType reflect.Type) error {
	arr, ok := c.Data[key].([]interface{})
	if !ok {
		return errors.New("key does not exist or value is not an array")
	}

	// Marshal the value to JSON.
	_, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Find the index of the value to remove.
	index := -1
	for i, v := range arr {
		// Unmarshal the value in the array from JSON to the desired type.
		vValue := reflect.New(valueType).Interface()
		vJSON, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(vJSON, vValue)
		if err != nil {
			return err
		}
		// Compare the unmarshaled value to the original value.
		if reflect.DeepEqual(reflect.ValueOf(vValue).Elem().Interface(), value) {
			index = i
			break
		}
	}

	// If the value was not found, return an error.
	if index == -1 {
		return errors.New("value not found in array")
	}

	// Remove the value from the array.
	c.Data[key] = append(arr[:index], arr[index+1:]...)
	return nil
}

func (c *Config) Save() error {
	return saveConfig(c.FileName, c)
}

func saveConfig(filename string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(filename string, standardConfig bool) (*Config, error) {
	log.Println("Trying to load " + filename + " ...")
	if standardConfig {
		if !FileExists(filename) {
			log.Println("Config file not found, creating new one ...")
			err := newConfig(filename)
			if err != nil {
				return nil, err
			}
		}
	}
	log.Println("Loading config ...")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	config.FileName = filename
	log.Println("Config loaded successfully")
	return &config, nil
}

type ExampleConfig struct {
	Data interface{} `json:"data"`
}

type ExampleConfigIngredients struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	APIKey   string `json:"apiKey"`
	Database struct {
		Driver string `json:"driver"`
		DSN    string `json:"dsn"`
	} `json:"database"`
	Logging struct {
		Level string `json:"level"`
		File  string `json:"file"`
	} `json:"logging"`
}

// NewConfig create new config
func newConfig(filename string) error {
	exampleConfIngredients := ExampleConfigIngredients{
		Port:   8080,
		Host:   "http://localhost",
		APIKey: RandomString(32),
		Database: struct {
			Driver string `json:"driver"`
			DSN    string `json:"dsn"`
		}{
			Driver: "mysql",
			DSN:    "",
		},
		Logging: struct {
			Level string `json:"level"`
			File  string `json:"file"`
		}{
			Level: "debug",
			File:  "",
		},
	}
	exampleConf := ExampleConfig{
		Data: exampleConfIngredients,
	}
	data, err := json.MarshalIndent(exampleConf, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	return err
}

func NewCustomConfig(filename string, content interface{}) error {
	exampleConf := ExampleConfig{
		Data: content,
	}
	data, err := json.MarshalIndent(exampleConf, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	return err
}

// FileExists Exists file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// generate random alphanumeric string with length
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
