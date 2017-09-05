import React from 'react';
import Navbar from 'navbar';
import LoginOrRegister from 'login-register';
import TodoOutline from 'todo-outline';
import TasksList from 'taskslist';

export default class FunToDone extends React.Component {
  render() {
    return (
      <Navbar />
      <LoginOrRegister />
      <TodoOutline />
      <TasksList />
    );
  }
}
ReactDOM.render(< Index / >, document.getElementById('index'));
