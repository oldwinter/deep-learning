import React from "react";

import { Grid } from "@alifd/next";

import "./index.css";

const Row = Grid.Row;

const Col = Grid.Col;

class Aaa extends React.Component {
  constructor(props, context) {
    super(props);
  }

  render() {
    return (
      <div>
        <span>文本</span>
        <a href="https://f2e.alibaba-inc.com/" target="_blank">
          链接
        </a>
        <Row style={{}}>
          <Col>
            <span>第一列</span>
          </Col>
          <Col>
            <span>第二列</span>
          </Col>
          <Col>
            <span>第三列</span>
          </Col>
        </Row>
      </div>
    );
  }
}

export default Aaa;
