<template>
  <router-view/>
</template>

<style lang="scss">
  *{
    margin: 0;
    padding: 0;
  }
  body,html{
    margin: 0;
    padding: 0;
  }
  .el-table--enable-row-hover .el-table__body tr:hover {
  background: rgb(184, 224, 250);
  border: 1px solid #313463;
}
/*表格鼠标悬停的样式（背景颜色）*/
.el-table tbody tr:hover > td {
  background-color: #2cce34 !important;
}


</style>