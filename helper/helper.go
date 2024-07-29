package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"` // insialisasi meta dan ambil meta dari struct meta
	Data interface{} `json:"data"` // inisialisasi data agar bisa dimasukkan berjenis data apapun alias dinamis
}

type Meta struct {
	// inisialisasi dalam struct meta
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// fungsi bertipe return response
func APIResponse(message string, code int, status string, data interface{}) Response {

	// instansiasi meta dari struct meta
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	// instansiasi data dari struct response dan meta dari struct meta
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	// return si response
	return jsonResponse
}

// Fungsi format validation error
func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

/*
 *	untuk membuat struktur heper seperti:
 *  Penjelasan: jika awalnya hurus kecil maka akan dianggap Private dan jika awalnya huruf kapital maka akan dianggap public
 *
 *	meta: {
 *		message: 'Your account has been created',
 *		code: 200,
 *		status: 'success'
 *	},
 *	data: {
 *		id: 1,
 *		name: "Agung Setiawan",
 *		occupation: "content creator",
 *		email: "com.agungsetiawan@gmail.com",
 *		token: "peterpanyangterdalam"
 * 	}
 *
 *	// penjelasan
 *	jika data sudah berhasil dimasukkan alias success dengan code 200
 *	maka tampilkan juga meta dengan pesan "Your account has been created"
 *
 */
