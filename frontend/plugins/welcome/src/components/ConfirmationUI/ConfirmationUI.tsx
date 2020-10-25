import React, { FC, useEffect,useState } from 'react';
import {Link as RouterLink} from 'react-router-dom'
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
} from '@material-ui/core';
import ComponentsTable from '../Table/Tables'
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntBorrow } from '../../api/models/EntBorrow';
import { EntStadium } from '../../api/models/EntStadium';
import { EntConfirmation } from '../../api/models/EntConfirmation';


const HeaderCustom = {
  minHeight: '50px',
};

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 20,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  table: {  
    minWidth: 650,
  },
}));

const Confirmation: FC<{}> = () =>{
  const UI = { giveName : 'Confirmation'}
  const User = { giveName : 'Woraphon Chutsungnoen'}
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  
  const [Confirmations, setConfirmations] = useState<EntConfirmation[]>([]);
  const [bookingstart, SetBookingstart] = useState(String);
  const [bookingend, SetBookingend] = useState(String);

  const [stadiums, setStadiums] = useState<EntStadium[]>([]);
  const [borrows, setBorrows] = useState<EntBorrow[]>([]);

  const [stadiumid, setStadiumID] = useState(Number);
  const [borrowid, setBorrowID] = useState(Number);

  useEffect(() =>{
    const getStadiums = async () =>{
      const res =  await api.listStadium({limit:10,offset:0});
      setLoading(false);
      setStadiums(res);
    }
    getStadiums();
    const getBorrows = async () =>{
      const res = await api.listBorrow({limit:10,offset:0});
      setLoading(false);
      setBorrows(res);
    }
    getBorrows();
    const getConfirmations = async () => {
      const res = await api.listConfirmation({ limit: 10, offset: 0 });
      setLoading(false);
      setConfirmations(res);
      console.log(res);
    };
    getConfirmations();
  },[loading]);

  const Bookingstarthandlechange = (event : any) =>{
    SetBookingstart(event.target.value as string);
  }

  const Bookingendhandlechange = (event : any) =>{
    SetBookingend(event.target.value as string);
  }

  const Stadiumhandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStadiumID(event.target.value as number);
  }
  
  const Borrowhandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBorrowID(event.target.value as number);
  }

  const CreateConfirmation = async ()=>{
    const confirmation = {
      user :  1,
      stadium : stadiumid,
      borrow : borrowid,
      bookingstart : bookingstart + ":00+07:00",
      bookingend : bookingend + ":00+07:00",
    };
    console.log(confirmation);
    const res: any = await api.createConfirmation({ confirmation : confirmation});
    if (res.id != '') {
      window.location.reload(false);
    }
  }
  console.log(Confirmations);

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`${UI.giveName}`}>
        <Avatar/>
        <div style={{ marginLeft: 10 }}>{User.giveName}</div>
      </Header>
      <Content>
        <ComponentsTable></ComponentsTable>
        <Grid item xs={12}>
        </Grid>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>สนามกีฬา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสนาม</InputLabel>
                <Select
                  name="stadium"
                  value={stadiumid || ""}
                  onChange={Stadiumhandlechange}
                >
                  {stadiums.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.namestadium}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ประเภท</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภท</InputLabel>
                <Select
                  name="borrow"
                  value={borrowid || ""}
                  onChange={Borrowhandlechange}
                >
                  {borrows.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.type}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>วันเวลาเริ่มต้นการจอง</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันและเวลา"
                  name="addday"
                  type="datetime-local"
                  value={bookingstart}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={Bookingstarthandlechange}
                />
              </form>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>วันเวลาสุดท้ายการจอง</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันและเวลา"
                  name="addday"
                  type="datetime-local"
                  value={bookingend}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={Bookingendhandlechange}
                />
              </form>
            </Grid>
            <Grid item xs={12}>
            <div className={classes.paper}>
              <Button

                variant="contained"
                color="primary"
                size="large"
                onClick={CreateConfirmation}
              >
                ยืนยันการจอง
              </Button>

              <Button

              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/StadiumUI"
              variant="contained"
              color="default"
              size="large"
              >
                ยกเลิก
              </Button>
              </div>
            </Grid>
          </Grid>
          <Grid item xs={10}></Grid>
        </Container>
      </Content>
    </Page>
  );


};
export default Confirmation;