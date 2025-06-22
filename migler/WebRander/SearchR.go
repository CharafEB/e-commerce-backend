package webrander

import (
	"fmt"
	"github/think.com/dots"
	"net/http"
	"strings"
)


// GetContentR renders the article content page

func ContentR(w http.ResponseWriter, res []dots.Products, queryValue string) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="ar" dir="rtl">
<head>
    <title>نتائج البحث:</title>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: 'Tajawal', Arial, sans-serif;
            padding: 20px;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
        }
        h1 {
            text-align: center;
            color: #2c3e50;
            margin-bottom: 30px;
        }
        .result-item {
            margin-bottom: 20px;
            padding: 20px;
            border: 1px solid #ddd;
            border-radius: 10px;
            background-color: #fff;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
            transition: transform 0.3s, box-shadow 0.3s;
        }
        .result-item:hover {
            transform: translateY(-5px);
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
        }
        .result-item strong {
            color: #1abc9c;
            font-size: 18px;
        }
        .result-item p {
            margin: 10px 0 0;
            line-height: 1.6;
            color: #555;
        }
        .back-button {
            display: inline-block;
            margin-top: 20px;
            padding: 10px 20px;
            background-color: #1abc9c;
            color: white;
            text-decoration: none;
            border-radius: 25px;
            font-size: 16px;
            transition: background-color 0.3s;
        }
        .back-button:hover {
            background-color: #16a085;
        }
    </style>
</head>
<body>
    <h1>نتائج البحث: %s</h1>
    %s
    <a href="#" class="back-button">العودة</a>
    <script>
        function makeRequest(val) {
            const searchURL = "http://localhost:3000/articlesContent/" + encodeURIComponent(val);
            window.open(searchURL, "_blank");
        }
    </script>
</body>
</html>
    `, queryValue, generateResultsHTMLForContent(res))
}

func generateResultsHTMLForContent(results []dots.Products) string {
	var html strings.Builder
	for index := range results {
		html.WriteString(fmt.Sprintf(`
        <div class="">
            <p><strong>العنوان:</strong> %s</p>
            <p><strong>المحتوى:</strong> %s</p>
        </div>
    `, results[index].ProductsName, results[index].ProductsShortDes))
	}
	return html.String()
}
