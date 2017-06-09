class LoginOrRegister extends React.Component {
  // TODO: convert to handleInputChange from here:
  // https://facebook.github.io/react/docs/forms.html
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      email: '',
      password: '',
      confpassword: '',
      remember: false
    };
    this.handleUsernameChange = this.handleUsernameChange.bind(this);
    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handleConfPasswordChange = this.handleConfPasswordChange.bind(this);
    this.handleRememberChange = this.handleRememberChange.bind(this);
  }
  handleUsernameChange(event) {
    this.setState({username: event.target.value});
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
                    <a href="#" className="active" id="login-form-link">Login</a>
                  </div>
                  <div className="col-xs-6">
                    <a href="#" id="register-form-link">Register</a>
                  </div>
                </div>
                <hr/>
                <div className="row">
                  <div className="col-lg-12"></div>
                    {/* TODO: use JS to POST data, not the default form handling */}
                  <form id="login-form" action="login" method="post" role="form">
                    <div className="form-group">
                      <input type="text" name="username" id="username" tabIndex="1" className="form-control" placeholder="Username" value={this.state.username} onChange={this.handleUsernameChange}/>
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
                  {/* TODO: use JS to POST data, not the default form handling */}
                  <form id="register-form" action="register" method="post" role="form">
                    <div className="form-group">
                      <input type="text" name="username" id="username" tabIndex="1" className="form-control" placeholder="Username" value={this.state.username} onChange={this.handleUsernameChange}/>
                    </div>
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
ReactDOM.render(< LoginOrRegister / >,
    document.getElementById('login-or-register'));
    /* TODO: figure out how to do this in React, not jQuery */
    $('#login-form-link').click(function(e) {
      $("#login-form").delay(100).fadeIn(100);
      $("#register-form").fadeOut(100);
      $('#register-form-link').removeClass('active');
      $(this).addClass('active');
      e.preventDefault();
    });
    $('#register-form-link').click(function(e) {
      $("#register-form").delay(100).fadeIn(100);
      $("#login-form").fadeOut(100);
      $('#login-form-link').removeClass('active');
      $(this).addClass('active');
      e.preventDefault();
    });
