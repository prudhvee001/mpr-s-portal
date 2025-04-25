import React from 'react';
import { Route, Switch } from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <h1>Welcome to the Student Portal</h1>
      <Switch>
        <Route path="/user" render={() => <h2>User Service</h2>} />
        <Route path="/results" render={() => <h2>Results Service</h2>} />
        <Route path="/notifications" render={() => <h2>Notifications Service</h2>} />
        <Route path="/info" render={() => <h2>Info Service</h2>} />
      </Switch>
    </div>
  );
}

export default App;
