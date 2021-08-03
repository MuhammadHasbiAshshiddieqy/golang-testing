package unittests

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
  tests := []struct {
    name      string
    request   string
    expected  string
  }{
    {
      name: "helloHasbi",
      request: "Hasbi",
      expected: "Hello, Hasbi!",
    },
    {
      name: "helloAyu",
      request: "Ayu",
      expected: "Hello, Ayu!",
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T)  {
      result := HelloWorld(test.request)
      assert.Equal(t, test.expected, result)
    })
  }
}
