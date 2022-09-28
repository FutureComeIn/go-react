import * as React from "react";

import PropTypes from "./js/prop-types"
import Axios from 'axios';

export default first;

function mapList() {
    const numbers = [1, 2, 3, 4, 5];
    return numbers.map((numbers,index) =>
                <li key={index}>
                    {numbers}-{index}
                </li>
           );
}

class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.state = {date: new Date(), username: "张安", lastGistUrl: "baidu.com", clickCount: 0};
        // 在dom中引用的方法，需要bind绑定
        this.handleClick = this.handleClick.bind(this);
    }
    // 生命周期钩子：当Clock的输出插入到DOM中时调用
    componentDidMount() {
        // Clock 组件要求浏览器设置一个定时器，每秒钟调用一次 tick()
        this.timerID = setInterval(
            () => this.tick(),
            1000
        );
        this.ajax();
    }
    // Clock被从DOM中移除时调用，定时器也就会被清除
    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    tick() {
        // 通过setState()，React知道状态已经改变，并再次调用render()来确定屏幕上应当显示什么
        this.setState({date: new Date()});
    }

    handleClick() {
        this.setState(function(state) {
            return {clickCount: state.clickCount + 1};
        });
    }

    ajax(){
        Axios.get("https://api.github.com/users/octocat/gists").then((response) => {
            let lastGist = response[0];
            this.setState({
                username: lastGist.owner.login,
                lastGistUrl: lastGist.html_url
            });
        });
    }
    render() {
        return (
            <div>
                <h1>Hello, {this.state.username}--{this.state.lastGistUrl}</h1>
                <h2>现在是 {this.state.date.toLocaleTimeString()}.</h2>
                <h2 onClick={this.handleClick}>点我！点击次数为: {this.state.clickCount}</h2>
                {mapList()}
            </div>
        );
    }
}

// 验证器，name属性要为String
Clock.propTypes = {
    name: PropTypes.string
};

function first() {
    return <Clock date={new Date()}/>
}

