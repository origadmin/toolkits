package rand

import "testing"

func TestGenerateRandom(t *testing.T) {
	length := 10
	random := GenerateRandom(length)
	t.Logf("Generated random: %s", random)
	if len(random) != length {
		t.Errorf("Expected salt length %d, but got %d", length, len(random))
	}

	length = 1024
	random = GenerateRandom(length)
	t.Logf("Generated random: %s", random)
	if len(random) != length {
		t.Errorf("Expected salt length %d, but got %d", length, len(random))
	}

}

func TestGenerateSalt(t *testing.T) {
	random := GenerateSalt()

	expectedLength := DefaultSaltSize
	t.Logf("Generated salt: %s", random)
	if len(random) != expectedLength {
		t.Errorf("Expected salt length %d, but got %d", expectedLength, len(random))
	}
}
