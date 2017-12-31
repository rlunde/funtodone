import React, { Component } from "react"
import { Navbar, Nav, NavItem, NavDropdown, MenuItem } from 'react-bootstrap';

class App extends Component {
  handleSelect(selectedKey) {
    alert(`selected ${selectedKey}`);
  }
  render() {
    return (
      <div>
        <header className="funtodone-header">
          <Navbar inverse className="funtodone-navbar">
            <Navbar.Header>
              <Navbar.Brand>
                <a href="/">FunToDone</a>
              </Navbar.Brand>
            </Navbar.Header>
            <Nav activeKey={1} onSelect={this.handleSelect}>
              <NavDropdown eventKey={1} title="Collection" id="basic-nav-dropdown">
                <MenuItem eventKey={1.1}>Load</MenuItem>
                <MenuItem eventKey={1.2}>Save</MenuItem>
                <MenuItem eventKey={1.3}>New</MenuItem>
                <MenuItem divider />
                <MenuItem eventKey={1.4}>Search</MenuItem>
              </NavDropdown>
            </Nav>
            <Nav pullRight>
              <NavItem eventKey={2} href="/home">Sign In</NavItem>
              <NavItem eventKey={3} title="Sign In">Sign Up</NavItem>
              <NavItem eventKey={4} disabled>Sign Out</NavItem>
            </Nav>
          </Navbar>
        </header>
        <div className="funtodone-body">
          <aside className="funtodone-aside">
          <div className="tasks-and-collections">
          <div className="collections">collections</div>
          <div className="tasks">tasks</div>
          </div>
          </aside>

          <main className="funtodone-content">Main Panel</main>
          <aside className="funtodone-aside">Score</aside>
        </div>
        <footer>Footer</footer>
      </div>
    )
  }
}
export default App
