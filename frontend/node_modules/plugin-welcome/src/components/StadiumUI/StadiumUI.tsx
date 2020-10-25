import React, { FC} from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import {Link as RouterLink} from 'react-router-dom'
import {
  Container,
    Grid,
  Avatar,
  Button,
} from '@material-ui/core';
import{
    ContentHeader,
    Link,
}from '@backstage/core';
import ComponentsTable from '../Table/Tables'
const HeaderCustom = {
  minHeight: '50px',
};

const Confirmation: FC<{}> = () =>{
    const UI = { giveName : 'Confirmation'}
    const User = { giveName : 'Woraphon Chutsungnoen'}

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`${UI.giveName}`}>
        <Avatar src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>{User.giveName}</div>
      </Header>
      <Content>
      <ContentHeader title="">
         <Link component={RouterLink} to="/ConfirmationUI">
           <Button variant="contained" color="primary">
             Add Confirmation
           </Button>
         </Link>
       </ContentHeader>
       <ComponentsTable></ComponentsTable>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );


};
export default Confirmation;