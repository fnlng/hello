import React, { Component } from 'react'
import todosData from './todosData.js'
import TodoItem from './TodoItem'

class TodoList extends Component {
    constructor() {
        super()
        this.state = {
            todos: todosData
        }
        this.handleChange = this.handleChange.bind(this)

    }

    handleChange(id) {
        this.setState(prevState => {
            const updatedTodoList = prevState.todos.map(todo => {
                if (todo.id === id) {
                    todo.comleted = !todo.comleted;
                }
                return todo
            })
            return {
                todos: updatedTodoList
            }
        })
    }

    render() {
        const todoComponents = this.state.todos.map(todo => {
            return (
                <TodoItem
                    todo={todo}
                    key={todo.id}
                    handleChange={this.handleChange}
                />    
            )
        })
    

        return (
            <div className="todo-list">
                {todoComponents}
            </div>
        )
    }
}

export default TodoList