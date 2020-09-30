import React, { Component } from 'react';
import PropTypes from 'prop-types';
import TransparentBackground from '../../shared/Backgrounds';
import HomeHeading from '../../shared/Headings';
import axios from 'axios';
import './App.css';
import WeatherMap from '../Weathermap/WeatherMap'
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      searchTerm: 'Hangzhou',
      isError: false,
      forecastResult: '',
    };

    

    // Bind class methods to object instances
    this.onSearchSubmit = this.onSearchSubmit.bind(this);
    this.onSearchChange = this.onSearchChange.bind(this);
    this.onDataRecieved = this.onDataRecieved.bind(this);
  }

  
  onSearchChange(e) {
    let locateMap = new Map([['Hangzhou', "Hangzhou"], ['Shanghai', "Shanghai"], ['Beijing', "Beijing"]]);

    this.setState({ searchTerm: locateMap.get(e.target.value) });
  }

  onDataRecieved(forecastResult) {
    const { searchTerm } = this.state;
    const { history } = this.props;
    if (forecastResult.cod === '404') {
      this.setState({ isError: true });
    } else {
      this.setState({ isError: false,forecastResult });
      history.push({
        pathname: `/forecast/${searchTerm}`,
        state:forecastResult
        
      })
      
    }
  }

  onSearchSubmit(e) {
    const { searchTerm } = this.state;
    if (searchTerm.length > 0) {

      let url = `data/weather?locate=${searchTerm}`
      let config = {
        baseURL: window.location.protocol + "//" + window.location.host
      }
        axios.get(url,config).then((forecastResult) => {
          this.setState({ isError: false });
          this.onDataRecieved(forecastResult.data)
        }).catch(error => {
            this.setState({ isError: true });
            console.log(error)
        })
    }
    e.preventDefault();
  }

  onClickDelete= () =>{
    this.setState({
      isError:false
    })
  }

  render() {
    let errorTip ={};
    if(this.state.isError){
      errorTip = <div class="notification is-danger">
          <button class="delete" onClick={this.onClickDelete}></button>
          The current forecast backend is not available
        </div>
      
    }else {
      errorTip = <div></div>;
    }
    return (
      <TransparentBackground
        className="column is-half
          is-offset-one-quarter  has-text-centered"
      >
        <HomeHeading className="title">Please Choose City</HomeHeading>
        <form onSubmit={this.onSearchSubmit} className="field has-addons has-addons-centered">
          <div className="control">
            <div className="field">
              <div className="control">
                <div className="select">
                  <select onChange={this.onSearchChange}>
                    <option>Hangzhou</option>
                    <option>Shanghai</option>
                    <option>Beijing</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
          <div className="control">
            <button type="submit" className="button is-primary">
              Search
            </button>
          </div>
        </form>
        <div className="field has-addons has-addons-centered">
        {errorTip}
        </div>
        <WeatherMap></WeatherMap>
        
      </TransparentBackground>
    );
  }
}

App.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func.isRequired,
  }).isRequired,
};

export default App;
