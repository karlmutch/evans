package gateway

import (
	"bytes"
	"testing"

	"github.com/jhump/protoreflect/dynamic"
	"github.com/ktr0731/evans/adapter/internal/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONFileInputter(t *testing.T) {
	env := testhelper.SetupEnv(t, "helloworld.proto", "helloworld", "Greeter")

	envMsg, err := env.Message("HelloRequest")
	require.NoError(t, err)

	jsonInput := `{
	"name": "ktr",
	"message": "hi"
}`

	msg := dynamic.NewMessage(envMsg.Desc)
	err = msg.TrySetField(msg.FindFieldDescriptorByName("name"), "ktr")
	require.NoError(t, err)
	err = msg.TrySetField(msg.FindFieldDescriptorByName("message"), "hi")
	require.NoError(t, err)

	in := bytes.NewReader([]byte(jsonInput))
	inputter := NewJSONFileInputter(in)
	actual, err := inputter.Input(envMsg.Desc)
	require.NoError(t, err)

	assert.Exactly(t, actual, msg)
}
