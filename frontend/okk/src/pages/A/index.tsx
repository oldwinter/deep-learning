import React, { Component } from 'react';
import ActionTable from './components/ActionTable';
import ExpandTable from './components/ExpandTable';

export default function () {
  return (
    <div className="A-page">
      {/* action bar table */}
      <ActionTable />
      {/* 可展开的表格 */}
      <ExpandTable />
    </div>
  );
}
