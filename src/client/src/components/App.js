import React, {Component} from "react"

class App extends Component {
  render() {
    return (
      <div>FunToDone
        <header>Header</header>
        <div className="funtodone-body">
          <main className="funtodone-content">Main Panel</main>
          <nav className="funtodone-nav">Nav Panel</nav>
          <aside className="funtodone-aside">Aside</aside>
        </div>
        <footer>Footer</footer>
      </div>
    )
  }
}
export default App
