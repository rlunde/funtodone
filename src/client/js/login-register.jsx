class LoginOrRegister extends React.Component {
  // TODO: convert to handleInputChange from here:
  // https://facebook.github.io/react/docs/forms.html
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: '',
      confpassword: '',
      remember: false
    };
    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handleConfPasswordChange = this.handleConfPasswordChange.bind(this);
    this.handleRememberChange = this.handleRememberChange.bind(this);
    this.handleLoginFormClick = this.handleLoginFormClick.bind(this);
    this.handleRegisterFormClick = this.handleRegisterFormClick.bind(this);
  }
  handleEmailChange(event) {
    this.setState({email: event.target.value});
  }
  handlePasswordChange(event) {
    this.setState({password: event.target.value});
  }
  handleConfPasswordChange(event) {
    this.setState({confpassword: event.target.value});
  }
  handleRememberChange(event) {
    this.setState({remember: event.target.checked});
  }
  handleLoginFormClick(event) {
    event.preventDefault();
    var link = $('#login-form-link');
    $("#login-form").delay(100).fadeIn(100);
    $("#register-form").fadeOut(100);
    $('#register-form-link').removeClass('active');
    link.addClass('active');
  }
  handleRegisterFormClick(event) {
    event.preventDefault();
    var link = $('#register-form-link');
    $("#register-form").delay(100).fadeIn(100);
    $("#login-form").fadeOut(100);
    $('#login-form-link').removeClass('active');
    link.addClass('active');
  }
  register(url) {
    fetch(url, {
      method: "POST",
      body: JSON.stringify(this.state),
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      mode: 'no-cors'
    }).then(result => result.json()).then(items => {
      console.log(items);
      this.setState({items})
    });
  }
  login(url) {
    fetch(url, {
      method: "POST",
      body: JSON.stringify(this.state),
      headers: {
        "Content-Type": "application/json"
      },
      mode: 'no-cors'
    }).then(result => result.json()).then(items => {
      console.log(items);
      this.setState({items})
    });
  }
  handleSubmit(event) {
    event.preventDefault();
    console.log("submit from: " + event.target.id);
    var formdata = $("#" + event.target.id).serialize();
    console.log("data:" + formdata);
    $('#lrModal').modal('hide');
    /*TODO: redirect to error page or main page on result of POST to REST API*/
    /*TODO: get the base URL from configuration */
    const baseUrl = "http://localhost:8080";
    if (event.target.id === "register-form") {
      this.register(baseUrl + "/register");
    } else if (event.target.id === "login-form") {
      this.login(baseUrl + "/login");
    }
  }
  render() {
    return (
      <div className="modal fade" id="lrModal" role="dialog">
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              <button type="button" className="close" data-dismiss="modal">&times;</button>
            </div>
            <div className="modal-body">
              <div className="panel panel-login">
                <div className="panel-heading">
                  <div className="row">
                    <div className="col-xs-6">
                      <a href="#" className="active" id="login-form-link" onClick={this.handleLoginFormClick}>Login</a>
                    </div>
                    <div className="col-xs-6">
                      <a href="#" id="register-form-link" onClick={this.handleRegisterFormClick}>Register</a>
                    </div>
                  </div>
                  <hr/>
                  <div className="row">
                    <div className="col-lg-12"></div>
                    {/* TODO: use JS to POST data, not the default form handling
                        using no action, but instead using: onSubmit={this.handleSubmit}
                    */}
                    <form id="login-form" action="" onSubmit={this.handleSubmit.bind(this)} method="post" role="form">
                      <div className="form-group">
                        <input type="text" name="email" id="email" tabIndex="1" className="form-control" placeholder="Email" value={this.state.email} onChange={this.handleEmailChange}/>
                      </div>
                      <div className="form-group">
                        <input type="password" name="password" id="password" tabIndex="2" className="form-control" placeholder="Password" value={this.state.password} onChange={this.handlePasswordChange}/>
                      </div>
                      <div className="form-group text-center">
                        <input type="checkbox" tabIndex="3" className="" name="remember" id="remember" value={this.state.remember} onChange={this.handleRememberChange}/>
                        <label htmlFor="remember">
                          Remember Me</label>
                      </div>
                      <div className="form-group">
                        <div className="row">
                          <div className="col-sm-6 col-sm-offset-3">
                            <input type="submit" name="login-submit" id="login-submit" tabIndex="4" className="form-control btn btn-login" value="Log In"/>
                          </div>
                        </div>
                      </div>
                      <div className="form-group">
                        <div className="row">
                          <div className="col-lg-12">
                            <div className="text-center">
                              <a href="recover" tabIndex="5" className="forgot-password">Forgot Password?</a>
                            </div>
                          </div>
                        </div>
                      </div>
                    </form>
                    {/* TODO: use JS to POST data, not the default form handling
                        using no action, but instead using: onSubmit={this.handleSubmit}
                    */}
                    <form id="register-form" action="" onSubmit={this.handleSubmit.bind(this)} method="post" role="form">
                      <div className="form-group">
                        <input type="email" name="email" id="email" tabIndex="1" className="form-control" placeholder="Email Address" value={this.state.email} onChange={this.handleEmailChange}/>
                      </div>
                      <div className="form-group">
                        <input type="password" name="password" id="password" tabIndex="2" className="form-control" placeholder="Password" value={this.state.password} onChange={this.handlePasswordChange}/>
                      </div>
                      <div className="form-group">
                        <input type="password" name="confirm-password" id="confirm-password" tabIndex="2" className="form-control" placeholder="Confirm Password" value={this.state.confpassword} onChange={this.handleConfPasswordChange}/>
                      </div>
                      <div className="form-group">
                        <div className="row">
                          <div className="col-sm-6 col-sm-offset-3">
                            <input type="submit" name="register-submit" id="register-submit" tabIndex="4" className="form-control btn btn-register" value="Register Now"/>
                          </div>
                        </div>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-default" data-dismiss="modal">Close</button>
            </div>
          </div>
        </div>
      </div>
    );
  }
};
ReactDOM.render(< LoginOrRegister / >, document.getElementById('login-or-register'));
