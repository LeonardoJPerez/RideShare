package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// TryDecodeBase64 decodes a potential Base63 encoded string.
// Returns true if successful, along with decoded string.
// Returns false if string is not Base64, along with an empty string.
func TryDecodeBase64(s string) (string, bool) {
	output, err := base64.StdEncoding.DecodeString(s)
	return string(output), err == nil
}

// GetEnvVariable attempts to retrieve an ENV variable.
// Returns default value if no evn exist or value is empty.
func GetEnvVariable(varName string, defaultValue ...string) string {
	value := os.Getenv(varName)
	if value == "" {
		return strings.Join(defaultValue, "")
	}
	return value
}

// PrettyPrint prints a JSON representation in a pretty printed manner.
func PrettyPrint(item interface{}) {
	b, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	strBuff := out.String()
	fmt.Printf("[%v] - %s\n", strings.ToUpper(os.Getenv("ENV")), strBuff)
}
