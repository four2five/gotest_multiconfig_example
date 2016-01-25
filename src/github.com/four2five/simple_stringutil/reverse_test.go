package simple_stringutil

import (
    //"fmt"
    "github.com/koding/multiconfig"
    . "github.com/four2five/simple_config"
    "testing"
)

// Perform some basic test setup, like configuration parsing
func SetupTest(t *testing.T) *StringMetadata {
    // Create a new DefaultLoader
    m := multiconfig.New()
    // Create an empty struct for my configuration
    stringMetadata := new(StringMetadata)
    // Load the configuration data
    m.MustLoad(stringMetadata)

    return stringMetadata
}

func TestReverse(t *testing.T) {
    stringMetadataPtr := SetupTest(t)
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }

    if stringMetadataPtr.Encoding != "Binary" {
        t.Error("Incorrect encoding, expected the default \"Binary\", found: %s\n",
                 stringMetadataPtr.Encoding);
    }
}
