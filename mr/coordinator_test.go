package mr

import "testing"

func TestCoordinator(t *testing.T) {
	files := []string{
		"../main/words/pg-being_ernest.txt",
		"../main/words/pg-dorian_gray.txt",
		"../main/words/pg-frankenstein.txt",
	}
	_ = MakeCoordinator(files, 10)

}
