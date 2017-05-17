var TodoOutline = React.createClass({
    render: function() {
        return (
          <div id="todo_outline">
          <span>To Do:</span>
          <ol>
              <li>design the basic layout of the FunToDone single page app</li>
              <ol type="A">
                  <li>figure out how to add <a href="https://github.com/go-authboss/authboss">authboss</a> to do login/register</li>
                  <ol>
                      <li>User registration</li>
                      <li>Confirming registrations via e-mail</li>
                      <li>User authentication via password</li>
                      <li>Password recovery via e-mail</li>
                      <li>Reset a User's password</li>
                  </ol>
                  <li>generate a sample stack and sample list for each new user</li>
                  <li>show a list of compound lists (stacks, lists) belonging to user</li>
                  <li>create a new stack</li>
                  <li>
                      <ol type="1">
                          <li>add a task to top of stack</li>
                          <li>split a task</li>
                          <li>add a subtask</li>
                          <li>choose a task/subtask to work on</li>
                          <li>edit/complete/delete a task</li>
                      </ol>
                  </li>
              </ol>
          </ol>
        </div>
        );
    }
});
ReactDOM.render(< TodoOutline / >, document.getElementById('todo'));
