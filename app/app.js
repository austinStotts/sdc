import React, { PureComponent } from 'react';
import Axios from 'axios';

class App extends PureComponent {
  constructor(props) {
    super(props)
    this.state = {
      id: 5000000,
      limit: 1,
      list: []
    }
    this.getStuff = this.getStuff.bind(this);
  }

  getStuff () {
    Axios.get(`http://localhost:3000/database?id=${this.state.id}&limit=${this.state.limit}`)
    .then(result => {
      this.setState({list: result.data})
    })
  }

  render () {
    return (
      <div>
        <div className="whole-input-wrap">
          <div className="input-wrap">
            <label className="label">id:</label>
            <input className="input" type="text" onChange={(e => {this.setState({id: e.currentTarget.value})})}></input>
          </div>
          <div>
            <label className="label">amount:</label>
            <input className="input" type="text" onChange={(e => {this.setState({limit: e.currentTarget.value})})}></input>
          </div>
        </div>
        <a className="anchor" onClick={this.getStuff}>get the data</a>
        <div>
          {this.state.list.length === 0 ? null : this.state.list.map((item) => (
            <div className="stuff-wrapper" key={Math.random()}>
              <h4 className="id">{item.ID}</h4>
              <h4 className="user">{item.Username}</h4>
              <h4 className="text">{item.Text}</h4>
              <h4 className="created">{item.Created}</h4>
              <h4 className="project">{item.ProjectID}</h4>
            </div>
          ))}
        </div>
      </div>
    )
  }
}

export default App;
