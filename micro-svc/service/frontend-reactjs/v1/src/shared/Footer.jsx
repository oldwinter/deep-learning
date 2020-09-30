import React, { Component } from 'react';
import styled from 'styled-components';
import airCleanImg from '../img/ad-images/air-clean.jpg'
import tshirtImg from '../img/ad-images/t-shirt.jpg'
import mountainClimbingImg from '../img/ad-images/mountain-climbing.jpg'
import axios from 'axios';

const StyledFooterContainer = styled.footer`
  background-color: #AABE3D;
`;
const ImageContainer = styled.figure`
  justify-content: center;
  align-items: center;
  margin: 2em auto 1em;
`;

class Footer extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isBackEndOK: false,
      adImgName: '',
      img: ""
    };

    this.updateAd = this.updateAd.bind(this);
  }

  updateAd(r) {
    this.setState({ adImgName: r.adImgName })
    if (this.state.adImgName === "airCleanImg") {
      this.setState({ img: airCleanImg })
    } else if (this.state.adImgName === "tshirtImg") {
      this.setState({ img: tshirtImg })
    } else if (this.state.adImgName === "mountainClimbingImg") {
      this.setState({ img: mountainClimbingImg })
    }
  }

  componentWillMount() {
    let url = "/advertisement/ad"
    let config = {
      baseURL: window.location.protocol + "//" + window.location.host
    }
    axios.get(url, config).then((response) => {
      this.setState({
        isBackEndOK: true
      })
      this.updateAd(response.data)
    }).catch(error => {
      this.setState({
        isBackEndOK: false
      })
      console.log(error)
    })
  }

  render() {
    let imageContainers = {};
    if (this.state.isBackEndOK) {
      imageContainers = <ImageContainer className="footer-image-middle">
        <img src={this.state.img} alt="" />
      </ImageContainer>
    } else {
      imageContainers = <ImageContainer className="footer-image-middle">
        <div className="notification is-danger">
          <button className="delete"></button>
          The current advertising backend is not available
        </div>
      </ImageContainer>
    }
    return (
      <StyledFooterContainer className="footer-container column 
 has-text-centered">

        {imageContainers}
      </StyledFooterContainer>
    )
  }
}




export default Footer;
