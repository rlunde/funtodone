export default class TasksList extends React.component({
    render: function() {
        return (
            <div className="tasksList">
                <button className="placeholder">Just a placeholder - nothing working yet</button>
            </div>
        );
    }
});
ReactDOM.render(< TasksList / >, document.getElementById('tasks'));
