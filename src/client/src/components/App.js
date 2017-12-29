import React, { Component } from "react"
import { Navbar, Nav, NavItem } from 'react-bootstrap';

class App extends Component {
  handleSelect(selectedKey) {
    alert(`selected ${selectedKey}`);
  }
  render() {
    return (
      <div>
        <header className="funtodone-header">
          <Navbar>
            <Navbar.Header>
              <Navbar.Brand>
                <a href="/">FunToDone</a>
              </Navbar.Brand>
            </Navbar.Header>
            <Nav bsStyle="pills" activeKey={1} onSelect={this.handleSelect}>
              <NavItem eventKey={1} href="/home">NavItem 1 content</NavItem>
              <NavItem eventKey={2} title="Item">NavItem 2 content</NavItem>
              <NavItem eventKey={3} disabled>NavItem 3 content</NavItem>
            </Nav>
          </Navbar>
        </header>
        <div className="funtodone-body">
          <aside className="funtodone-aside">Aside</aside>
          <main className="funtodone-content">Main Panel</main>
        </div>
        <footer>Footer</footer>
      </div>
    )
  }
}
export default App
