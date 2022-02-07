import "@babel/polyfill";
import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route,Switch,Redirect  } from 'react-router-dom';
//用bulma实现布局和控件
import 'bulma/css/bulma.css';
//用fontawesome的开源图标库
// import "@fortawesome/fontawesome-free/css/all.min.css"

import App from './components/App';
import Forecast from './components/Forecast';
import Maintenanced from './components/Maintenanced';

import Sky, { SkyLandscape } from './shared/Sky';
import Footer from './shared/Footer';
import Header from './shared/Header';

const router = (
  <Router>
    <div>
      <Header />
      <SkyLandscape>
        <Sky />
        <Switch>
          <Route exact path="/dashboard" component={App} />
          <Route exact path="/forecast/:city" component={Forecast} />
          <Route exact path="/maintenanced" component={Maintenanced} />
          <Redirect to="/dashboard" />
          {/* <Route exact path="/map" component={WeatherMap}></Route> */}
        </Switch>
      </SkyLandscape>
      <Footer />
      
    </div>
  </Router>
);

ReactDOM.render(router, document.getElementById('root'));

if (module.hot) {
  module.hot.accept();
}
