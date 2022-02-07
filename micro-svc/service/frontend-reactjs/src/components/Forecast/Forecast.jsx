import React, { Component } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import icons from '../../shared/Icons';
import TransparentBackground from '../../shared/Backgrounds';
import axios from 'axios';

const Loader = styled.a`
  font-size: 3rem;
  border: 0;
`;

const ImageContainer = styled.figure`
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2em auto 1em;
`;

const Card = styled.div`
  transition: all 0.2s ease-in-out;
  background-color: rgba(255,255,255,0.7);
  &:hover {
    box-shadow: 0 4px 3px rgba(10, 10, 10, 0.1), 0 0 0 1px rgba(10, 10, 10, 0.1);
    transform: scale(1.1);
  }
`;

const Days = ({ forecastResult, match }) => (
  <div className="columns">
    {forecastResult.list.map((days, i) => {
      const formattedDate = days.dt.replace(/^\S/, function (s) { return s.toUpperCase(); });
      let [dressTemp, sportTemp] = days.recommendation ? days.recommendation.split(';') : "::"
      let dress = dressTemp.split(":")[1];
      let sport = sportTemp.split(":")[1];
      let RecommendationCards = <div></div>;

      //v2版本有推荐信息
      if (days.recommendation) {
        RecommendationCards = <Card className="card Istio-mgTop-md">
          <div className="card-content">
            <p className="subtitle is-4">{dress}</p>
            <p className="title is-2">{sport}</p>
          </div>
        </Card>
      }
      return (
        <div key={days.dt} className="column">
          <Card className="card">
            <div className="card-content">
              <p className="subtitle is-4">{formattedDate}</p>
              <p className="title is-2">{Math.round(days.temp.day)} ºC</p>
              <ImageContainer className="image is-128x128">
                <img src={icons[`icon${days.weather[0].icon}`]} alt="" />
              </ImageContainer>
            </div>
          </Card>
          {RecommendationCards}
        </div>
      );
    })}
  </div>
);

class Forecast extends Component {
  constructor(props) {
    super(props);
    this.state = {
      forecastResult: '',
      isLoading: true,
    };

    this.updateResult = this.updateResult.bind(this);
  }

  componentWillMount() {
    //从上个页面的搜索按钮跳转过来，则无需再请求接口获取数据
    if (this.props.history.location.state) {
      this.updateResult(this.props.history.location.state);
    } else {
      let url = `data/weather?locate=${this.props.location.pathname.split('/')[2]}`
      let config = {
        baseURL: window.location.protocol + "//" + window.location.host
      }
      axios.get(url, config).then((forecastResult) => {

        this.updateResult(forecastResult.data)

      }).catch(error => {
        console.log(error)
      })
    }
  }

  updateResult(state) {
    this.setState({
      forecastResult: state,
      isLoading: false,
    });
  }
 

  render() {
    const { forecastResult, isLoading } = this.state;
    const { match } = this.props;
    return (
      <TransparentBackground
        className="column is-10 is-offset-1 has-text-centered"
      >
        {isLoading
          ? <Loader className="button is-loading">
            Button
            </Loader>
          : <div>
            <h1 className="title">{forecastResult.city.name} Forcast</h1>
            <Days match={match} forecastResult={forecastResult} />
          </div>}
      </TransparentBackground>
    );
  }
}

Forecast.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
    params: PropTypes.object,
  }).isRequired,
  history: PropTypes.shape({
    location: PropTypes.object,
    state: PropTypes.object,
  }).isRequired,
};

Days.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
  forecastResult: PropTypes.shape({}).isRequired,
};

export default Forecast;
