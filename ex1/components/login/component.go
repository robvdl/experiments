package login

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

// Template is the template for the login form.
const Template = `<div>
  <div>Username: {{username}}</div>
  <input v-model="username" />
  <button @click="Login">Login</button>
</div>`

// LoginForm is a login form component.
type LoginForm struct {
	*js.Object

	Username string `js:"username"`
	Password string `js:"password"`
}

// Login is called when the login button is clicked.
func (l *LoginForm) Login() {
	println("hello", l.Username)
}

// New constructs the login form component.
func New() interface{} {
	return &LoginForm{
		Object: js.Global.Get("Object").New(),
	}
}

// Register registers the component with vue.js.
func Register() {
	vue.NewComponent(New, Template).Register("login-form")
}
