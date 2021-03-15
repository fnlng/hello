import React from "react"

function TodoItem(props) {
    const completedStyle = {
        fontStyle: 'italic',
        textDecoration: 'line-through'
    }

    const uncompletedStyle = {
        //
    }

    return (
        <div
            className="todo-item"
            onClick={() => props.handleChange(props.todo.id)}
        >
            <input
                onChange={() => {}}
                name="completed"
                type="checkbox"
            />
            <p style={props.todo.completed ? completedStyle : uncompletedStyle}>{props.todo.text}</p>
        </div>
    )
}

export default TodoItem