import React from "react";

import { Grid, Button, Search, Upload } from "@alifd/next";

import "./index.css";

const Row = Grid.Row;

const Col = Grid.Col;

class Aaaa extends React.Component {
  constructor(props, context) {
    super(props);
  }

  render() {
    return (
      <div>
        <Row style={{}}>
          <Col>
            <span>第一列</span>
          </Col>
          <Col>
            <span>第二列</span>
          </Col>
        </Row>
        <Button type="secondary">确定</Button>
        <Button type="secondary">确定</Button>
        <Search type="primary" />
        <Upload
          action="https://www.easy-mock.com/mock/5b713974309d0d7d107a74a3/alifd/upload"
          multiple={true}
          style={{ display: "inline-block" }}
          shape="card"
          useDataURL={false}
        >
          上传图片
        </Upload>
      </div>
    );
  }
}

export default Aaaa;
