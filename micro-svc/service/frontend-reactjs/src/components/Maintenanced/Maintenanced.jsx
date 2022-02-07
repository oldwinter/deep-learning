import React, { Component } from 'react';
import TransparentBackground from '../../shared/Backgrounds';
import styled from 'styled-components';
import MaintenancedImg from '../../img/500.svg';
const ImageContainer = styled.figure`
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2em auto 1em;
`;
class Maintenanced extends Component {
    constructor(props){
        super(props)
    }
    render(){
        return(
            <TransparentBackground
                className="has-text-centered"
            >
            <section className="hero is-dark is-medium">
                <div className="hero-body">
                    <div className="container">
                    <h1 className="title font-size-veryLarge">
                        500
                    </h1>
                    <h2 className="subtitle">
                        The server is in the maintenance
                    </h2>
                    </div>
                </div>
                </section>
            
            </TransparentBackground>
        )
    }
}
export default Maintenanced 