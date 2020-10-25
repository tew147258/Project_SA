import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntConfirmation} from '../../api/models/EntConfirmation';
import moment from 'moment'


const useStyles = makeStyles({
    table: {
      minWidth: 650,
    },
});
export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [Confirmations, setConfirmations] = useState<EntConfirmation[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getConfirmations = async () => {
          const res = await api.listConfirmation({ limit: 10, offset: 0 });
          setLoading(false);
          setConfirmations(res);
          console.log(res);
        };
        getConfirmations();
    }, [loading]);
    const deleteConfirmation = async (id: number) => {
      const res = await api.deleteConfirmation({ id: id });
      setLoading(true);
    };
   

    console.log(Confirmations);
    return (
      <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">Users</TableCell>
            <TableCell align="center">Stadiums</TableCell>
            <TableCell align="center">Borrow</TableCell>
            <TableCell align="center">Adddate</TableCell>
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
              <TableCell align="center">{moment(item.adddate).format("DD/MM/YYYY HH.mm น.")}</TableCell>
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
      );
     


}