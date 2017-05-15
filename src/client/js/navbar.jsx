var Navbar = React.createClass({
    render: function() {
        return (
          <div id="navbar">
            <nav className="navbar navbar-default navbar-fixed-top">
                <div className="container-fluid">
                    <div className="navbar-header">
                        <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                          <span className="sr-only">Toggle navigation</span>
                          <span className="icon-bar"></span>
                          <span className="icon-bar"></span>
                          <span className="icon-bar"></span>
                        </button>
                        <a className="navbar-brand" href="#">FunToDone</a>
                    </div>
                    <div className="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                        <ul className="nav navbar-nav">
                            <li className="dropdown">
                                <a href="#" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Task Collections <span className="caret"></span></a>
                                <ul className="dropdown-menu">
                                    <li><a href="#">Load a task collection</a></li>
                                    <li><a href="#">Start a new task collection</a></li>
                                    <li><a href="#">Export a task collection</a></li>
                                    <li><a href="#">Import a task collection</a></li>
                                    <li role="separator" className="divider"></li>
                                    <li><a href="#">Statistics / History</a></li>
                                    <li role="separator" className="divider"></li>
                                    <li><a href="#">Help</a></li>
                                </ul>
                            </li>
                        </ul>
                        <form className="navbar-form navbar-left">
                            <div className="form-group">
                                <input type="text" className="form-control" placeholder="Search for a task"/>
                            </div>
                            <button type="submit" className="btn btn-default">Submit</button>
                        </form>
                        <ul className="nav navbar-nav navbar-right">
                            <li><a href="#" data-toggle="modal" data-target="#lrModal">Login / Register</a></li>
                        </ul>
                    </div>
                </div>
            </nav>
          </div>
        );
    }
});
ReactDOM.render(< Navbar / >, document.getElementById('navdiv'));
