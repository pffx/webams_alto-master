import React, { useState, useRef } from 'react';
import PropTypes from 'prop-types';
// import { getContext } from '@nokia-csf-uxr/ccfk/common';
import SearchwChips from '@nokia-csf-uxr/ccfk/SearchwChips';
import _mapValues from 'lodash/mapValues';
import DataGrid from '@nokia-csf-uxr/ccfk/DataGrid';
import { AllCommunityModules } from '@ag-grid-community/all-modules';
import Button from '../button';
// import gridSearchWChipsColumnDefs from '@nokia-csf-uxr/ccfk/DataGrid/stories/utils/columnDefs/gridSearchWChipsColumnDefs';

// import data from '@nokia-csf-uxr/ccfk/DataGrid/stories/utils/data.json';

const DataGridTable = (props) => {
  const {
    dataList,
    gridColumnDefs,
    hasSearch,
    pageSize,
    hasExport,
    style,
  } = props;
  const [gridApi, setGridApi] = useState({});
  const [filterValue, setFilterValue] = useState([
    {
      type: 'field',
      id: 'field-1',
      value: ''
    }
  ]);
  const filterValueRef = useRef([]);
  const columnDefs = gridColumnDefs;
  const onGridReady = (params) => {
    setGridApi(params.api);
  };
  const externalFilterChanged = (newValue) => {
    filterValueRef.current = newValue;
    setFilterValue(newValue);
    gridApi && gridApi.onFilterChanged && gridApi.onFilterChanged();
  };

  const showOrHideOverlay = () => {
    console.log("showOrHideOverlay   gridApi = ",gridApi)
    if (gridApi && gridApi.getModel().rowsToDisplay) {
      if (gridApi.getModel().rowsToDisplay.length === 0) {
        gridApi.showNoRowsOverlay();
      } else {
        gridApi.hideOverlay();
      }
    }
  };

  const onFilterChanged = () => {
    showOrHideOverlay();
  };

  const isExternalFilterPresent = () => filterValueRef.current.filter(item => item.type === 'chip').length > 0;

  const doesExternalFilterPass = ({ data }) => {
    // console.log("doesExternalFilterPass   data = ",data)
    let pass = true;
    if (!isExternalFilterPresent()) {
      return pass;
    }
    const terms = filterValueRef.current
      .filter(chipsItem => chipsItem.type === 'chip')
        .map(chip => chip.value.toString().toLowerCase());

    terms.forEach(term => {
      let isTermMatchedWithACol = false;
      _mapValues(data, columnValue => {
        const stringColValue = (columnValue === undefined || columnValue === null) ? "" : columnValue.toString().toLowerCase()
        isTermMatchedWithACol = isTermMatchedWithACol || stringColValue.indexOf(term) > -1;
      });
      pass = pass && isTermMatchedWithACol;
    });
    return pass;
  };
  const onBtnCSVExport = () =>  {
    gridApi && gridApi.exportDataAsCsv();
  };
  const onBtnExcelExport = () =>  {
    // gridApi && console.log("zp  onBtnExcelExport gridApi = ",gridApi.getDataAsExcel())
     // /ag-grid-community does not support export to Excel
    gridApi && gridApi.exportDataAsExcel();
  };
  const onBtnPrint= () =>  {
    // gridApi && console.log("zp  onBtnPrint gridApi = ",gridApi)
    // gridApi && gridApi.exportDataAsExcel();
    window.print()
  };

  return (
    <>
      {/* {hasSearch&&<SearchwChips style={{marginLeft:"0rem",marginBottom:"0.5rem"}} data={filterValue} onChange={externalFilterChanged} size="small" />} */}
      
      <div className='row3'>
        <div className='row' style={{width:"30%"}}>
          {hasExport &&<Button style={{ width: "30%"}} onClick={onBtnCSVExport} title={"Export"}></Button>}
          {/* {hasExport &&<Button style={{ width: "30%"}} onClick={onBtnExcelExport}  title={"Excel"}></Button>} */}
          {hasExport &&<Button style={{ width: "30%"}} onClick={onBtnPrint}  title={"Print"}></Button>}
        </div>
        {hasSearch&&<SearchwChips style={{minWidth:"30%", maxWidth:"50%", marginBottom:"0.5rem"}} data={filterValue} onChange={externalFilterChanged} size="small" />}
      </div>
      <DataGrid
        modules={AllCommunityModules}
        wrapperProps={{
          style: {
            height: '100%',
            width: '100%',
            ...style
          }
        }}
        rowData={dataList}
        suppressContextMenu
        pagination={true}
        paginationAutoPageSize={true}
        paginationPageSize={pageSize}
        defaultColDef={{
          menuTabs: ['filterMenuTab']
        }}
        suppressCsvExport={false}
        // suppressExcelExport={false}
        columnDefs={columnDefs}
        onFilterChanged={onFilterChanged}
        onGridReady={onGridReady}
        isExternalFilterPresent={isExternalFilterPresent}
        doesExternalFilterPass={doesExternalFilterPass}
      />
    </>
  );
};
DataGridTable.propTypes = {
  dataList:PropTypes.array.isRequired,
  gridColumnDefs:PropTypes.array.isRequired,
  hasSearch:PropTypes.bool,
  hasExport:PropTypes.bool,
  pageSize:PropTypes.number,
  style: PropTypes.object,
};

DataGridTable.defaultProps = {
  hasSearch:true,
  hasExport:true,
  dataList:[],
  gridColumnDefs:[],
  pageSize:10,
  style:{},

}
export default DataGridTable;
