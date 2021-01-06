import React, { FC, useEffect,useState } from 'react';
import {Link as RouterLink} from 'react-router-dom';
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
}from '@material-ui/core';
import { EntBorrow } from '../../api/models/EntBorrow';
import { EntStadium } from '../../api/models/EntStadium';
import { DefaultApi } from '../../api/apis';
import { EntConfirmation} from '../../api/models/EntConfirmation';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import moment from 'moment'


const HeaderCustom = {
  minHeight: '50px',
};

// css style 
const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
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
}));

const Confirmation: FC<{}> = () =>{
  const UI = { giveName : 'Confirmation'}
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [loadingcon, setloadingcon] = useState(true);
  const [bookingstart, SetBookingstart] = useState(String);
  const [bookingend, SetBookingend] = useState(String);

  const [stadiums, setStadiums] = useState<EntStadium[]>([]);
  const [borrows, setBorrows] = useState<EntBorrow[]>([]);

  const [stadiumid, setStadiumID] = useState(1);
  const [borrowid, setBorrowID] = useState(Number);
let idstadium = Number(stadiumid)
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
    
  },[loading]);

  const [Confirmations, setConfirmations] = useState<EntConfirmation[]>([]);
  useEffect(() =>{
    const getConfirmations = async () => {
      const res = await api.getConfirmation({id : idstadium});
      setloadingcon(false)
      setConfirmations(res);
      console.log(res);
    };
    
    getConfirmations();
  },[loadingcon]);
   
  const Bookingstarthandlechange = (event : any) =>{
    SetBookingstart(event.target.value as string);
  }

  const Bookingendhandlechange = (event : any) =>{
    SetBookingend(event.target.value as string);
  }

  const Stadiumhandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStadiumID(event.target.value as number);
    setloadingcon(true)
  }
  
  const Borrowhandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBorrowID(event.target.value as number);
  }
  const deleteConfirmation = async (id: number) => {
    const res = await api.deleteConfirmation({ id: id });
    setLoading(true);
    window.location.reload(false);
  };

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
      </Header>
      <Content>
      <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">Users</TableCell>
            <TableCell align="center">Stadiums</TableCell>
            <TableCell align="center">Borrow</TableCell>
            <TableCell align="center">Bookingdate</TableCell>
            <TableCell align="center">Bookingstart</TableCell>
            <TableCell align="center">Bookingend</TableCell>
            <TableCell align="center">Hourstime</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          { Confirmations.map((item:any )=> (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.edges.confirmationUser.name}</TableCell>
              <TableCell align="center">{item.edges.confirmationStadium.namestadium}</TableCell>
              <TableCell align="center">{item.edges.confirmationBorrow.type}</TableCell>
              <TableCell align="center">{moment(item.bookingdate).format("DD/MM/YYYY HH.mm น.")}</TableCell>
              <TableCell align="center">{moment(item.bookingstart).format("DD/MM/YYYY HH.mm น.")}</TableCell>
              <TableCell align="center">{moment(item.bookingend).format("DD/MM/YYYY HH.mm น.")}</TableCell>
              <TableCell align="center">{item.hourstime}</TableCell>
              <Button
                onClick={() => {
                  deleteConfirmation(item.id);
                }}
                style={{ marginLeft: 10 }}
                variant="contained"
                color="secondary"
              >
                Delete
              </Button>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
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
                style={{ marginLeft: 50 }}
                component={RouterLink}
                to="/"
                variant="contained"
                color="secondary"
                size="large"
              >
                Logout
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