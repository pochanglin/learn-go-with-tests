package maps

import "testing"

func TestSearch(t *testing.T){

	dictionary := Dictionary{"Test": "Just a Test"}

	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("Test")
		want := "Just a Test"
		assertStrings(t,got,want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, got := dictionary.Search("Unknown")
		assertError(t,got,ErrNotFound)
	})

}

func TestAdd(t *testing.T){

	t.Run("New Word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := d.Add(word, definition)

		assertError(t,err,nil)
		assertDefinition(t, d, word, definition)
	})

	t.Run("Existing Word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T){
	
	t.Run("New Word", func(t *testing.T) {
		word := "test"
		definition := "Just a Test"
		newDefinition := "new definition"
		d := Dictionary{word:definition}

		err := d.Update(word,newDefinition)

		assertError(t,err,nil)
		assertDefinition(t,d,word,newDefinition)
	})

	t.Run("Existing Word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T)  {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}