# GoLang-Rest-API
Simple ReST API Using Go Lang

---

### What's Included
1. A Simple ReST API using Go Lang
2. Mux Router
3. Currently data managed from memory


### @Todo
1. Add Database Configuration
2. Make architecture for Big One

### How to start
**Step 1:**
Install Go language on your machine.

**Step 2:**
Clone this repository
```bash
git clone https://github.com/ManiruzzamanAkash/GoLang-Rest-API.git
```

**Step 3:**
Go to that project, if you're on another directory.

**Rename** the folder to something else -
```bash
goapi-test
```

Go to that folder -
```bash
cd goapi-test
```

**Step 4:**
Build App
```bash
go build
```

**Step 5:**
Run
```bash
./goapi-test
```

### Test the API Endpoints:**


**All Books**

URL: http://localhost:8000/api/books

Method: `get`


**Get Single Book**
URL: http://localhost:8000/api/books/1

Method: `get`


**Create a book -**

URL: http://localhost:8000/api/books

Method: `post`

```json
{
	"Title" : "Book Created",
	"Isbn": "819289823",
	"Author": {
		"FirstName": "Maniruzzaman",
		"LastName": "Akash"
	}
}
```


**Update a Book**

URL: http://localhost:8000/api/books/2

Method: `put`

Where `2` is the ID

```json
{
	"Title" : "Book Updated for 2",
	"Isbn": "819289823",
	"Author": {
		"FirstName": "Maniruzzaman",
		"LastName": "Akash"
	}
}
```


**Delete a Book**

URL: http://localhost:8000/api/books/31847

Method: `delete`

Where `31847` is the ID