# 📚 Book Search API with Golang & Typesense

This project is a **Book Search API** built with **Golang (Gin framework)** and **Typesense** as the search engine.  
It allows users to **search for books** by title or author name with **fast and accurate results**.

---

## 🚀 Features

✅ **Search Books** by title or author name  
✅ **Fast & Scalable Search** powered by Typesense  
✅ **Built with Golang & Gin Framework**  
✅ **JSON-based Data Storage**  
✅ **Easy Setup & Deployment**  

---

## 🛠️ Tech Stack

| Technology  | Description  |
|-------------|-------------|
| **Golang**  | Backend API Development  |
| **Gin**  | Web Framework for Golang  |
| **Typesense**  | Search Engine  |
| **JSONL**  | Data Storage Format  |

---

## ⚙️ Installation & Setup

### 1️⃣ Clone the Repository

First, download the project from GitHub using **Git**:

```sh
git clone git@github.com:abutahshin/Text-Search-Engine.git
cd Text-Search-Engine
```
### 2️⃣ Install Golang Dependencies

Make sure you have Golang installed on your machine. Then, install the required dependencies:
```sh
go mod tidy
```
### 3️⃣ Set Up Typesense
Typesense is required to store and search the book data.

**A. Install & Run Typesense Locally**
	•	Download & install Typesense from: Typesense Setup Guide:[Typesense Setup Guide](https://typesense.org/docs/guide/)
	•	Start a local Typesense instance:

```
typesense-server --data-dir ./typesense-data --api-key=YOUR_API_KEY
```

**B. Create a Collection in Typesense**
The books collection must be created before importing data.
Use cURL to create the collection:

***C. Import Book Data into Typesense***

Ensure you have a books.jsonl file in your project directory.
Run the following command to import the book data:

```
curl -X POST "http://your_typese_url/collections/books/documents/import?action=create" \
     -H "X-TYPESENSE-API-KEY: YOUR_API_KEY" \
     -H "Content-Type: text/plain" \
     --data-binary @books.jsonl
```

### 4️⃣ Run the API Server

After setting up Typesense, start the Golang API server:

```sh
go run main.go
```

By default, the API runs on:
```
http://localhost:8080.
```

### 🔍 Usage

**Search for Books**

Use this endpoint to search for books by title or author:
```
http://localhost:8080/search?q=harry
```
### 📌 Example API Response:

```
{
  "found": 2,
  "hits": [
    {
      "title": "Harry Potter and the Sorcerer's Stone",
      "authors": ["J.K. Rowling"],
      "average_rating": 4.5,
      "publication_year": 1997,
      "image_url": "https://example.com/image1.jpg",
      "ratings_count": 1200
    },
    {
      "title": "Harry Potter and the Chamber of Secrets",
      "authors": ["J.K. Rowling"],
      "average_rating": 4.6,
      "publication_year": 1998,
      "image_url": "https://example.com/image2.jpg",
      "ratings_count": 1100
    }
  ]
}
```

### 📖 API Endpoints:

|Method |Endpoint|Description|
|-------|-----------------|-------------------------------|
|**GET**|/search?q={query}|Search books by title or author|

### For Details About Typesense Api Endpoints:

[Explore This Blog](https://github.com/typesense/typesense-go)


### Postman Collection For Testing Typesense:

[Postman Collection](https://github.com/typesense/postman)