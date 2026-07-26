package progress

import "testing"

func TestLoadSave(t *testing.T) {
	path := t.TempDir() + "/progress.json"
	store, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	store.Complete("regex-01-literals/timeout")
	if err := store.Save(path); err != nil {
		t.Fatal(err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if !loaded.IsComplete("regex-01-literals/timeout") {
		t.Fatal("completed exercise was not persisted")
	}
}
