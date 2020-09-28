import React, { Component } from 'react';
import styled from 'styled-components';
import Cookies from 'js-cookie'
// import FormValidation from 'formvalidation/dist/js/formValidation.min.js'

const StyledHeaderContainer = styled.header`
  background-color: #65bffd;
`;

class LoginModal extends Component {
    constructor(props) {
        super(props);
        this.state = {
            modalState: false,
            username: '',
            pwd: ''
        }
        console.log(props)
    }
    render() {
        console.log(this)
        if (!this.props.modalState) {
            return null;
        }
        return (
            <div className="modal is-active">
                <div className="modal-background" onClick={this.props.closeModal} />
                <div className="modal-card">
                    <header className="modal-card-head">
                        <p className="modal-card-title">{this.props.title}</p>
                        <button className="delete" onClick={this.props.closeModal} />
                    </header>
                    <section className="modal-card-body">
                        <div className="content">
                            {this.props.children}
                        </div>
                    </section>
                    <footer className="modal-card-foot">
                        <button type="submit" className="button is-link" onClick={this.props.okAction}>Login</button>
                        <button className="button" onClick={this.props.closeModal}>Cancel</button>
                    </footer>
                </div>
            </div>
        );
    }

}
// LoginModal.propTypes = {
//   closeModal: React.PropTypes.func.isRequired,
//   modalState: React.PropTypes.bool.isRequired,
//   title: React.PropTypes.string
// }

class Header extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isLogin: false,
            modalState: false,
            username: "huawei_ASM",
            pwd: ""
        };
    }

    componentWillMount() {
        if (Cookies.get("user")) {
            this.setState({
                isLogin: true
            })

        } else {
            this.setState({ isLogin: false })
        }
    }
    onLogin = () => {

        this.setState({
            modalState: true
        })

        // Cookies.set('user', 'asm');
        // this.setState({
        //     isLogin: true
        // })


    }

    onLogout = () => {
        Cookies.remove('user');
        this.setState({
            isLogin: false
        })
    }
    toggleLoginAction = () => {

        Cookies.set('user', this.state.username);
        this.setState({
            isLogin: true,
            modalState: false

        })


    }
    toggleModal = () => {
        this.setState((prev, props) => {
            const newState = !prev.modalState;

            return { modalState: newState };
        });



    }
    onChangeHandler = (key, val) => {

        this.setState({
            [key]: val.target.value
        })
    }

    render() {

        let loginComponent =
            <nav className="level">
               
                    {this.state.isLogin ? (
                        <div className="level-left">
                            <div className="level-item button is-link" onClick={this.onLogout}>Logout</div>
                            <div className="level-item ">Current User: {Cookies.get('user')}</div>
                        </div>
                    )
                        :
                        (<div className="level-left">
                            <div className="button is-link" onClick={this.onLogin}>Login</div>
                        </div>)}
            
            </nav>


        return (<StyledHeaderContainer className="header-container">
            {loginComponent}
            <section className="section">
                <div className="container">
                    <LoginModal
                        okAction={this.toggleLoginAction}
                        closeModal={this.toggleModal}
                        modalState={this.state.modalState}
                        title="Login weather forecast"
                    >
                        <form id="loginForm">
                            <div class="field">
                                <label class="label">Username</label>
                                <p class="control has-icons-left has-icons-right">
                                    <input class="input" type="text" onChange={e => this.onChangeHandler('username', e)} placeholder="huawei_ASM" />
                                    <span class="icon is-small is-left">
                                        <i class="fas fa-envelope"></i>
                                    </span>
                                    <span class="icon is-small is-right">
                                        <i class="fas fa-check"></i>
                                    </span>
                                </p>
                                <p class="help">Click login button will generate cookie: "user={this.state.username}" in your browser</p>
                            </div>
                            {/* <div class="field">
                            <p class="control has-icons-left">
                                <label class="label">Password</label>
                                <input class="input" type="password" onChange={e =>this.onChangeHandler('pwd',e)} placeholder="Password" />
                                <span class="icon is-small is-left">
                                    <i class="fas fa-lock"></i>
                                </span>
                            </p>
                            </div> */}
                        </form>

                    </LoginModal>
                </div>
            </section>


        </StyledHeaderContainer>
        )
    }
}




export default Header;
