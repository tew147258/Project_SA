import React, { FC,useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import {Link as RouterLink} from 'react-router-dom'

function Copyright() {
    return (
      <Typography variant="body2" color="textSecondary" align="center">
        {'Copyright © '}
        Group 16 SA-63{' '}
        {new Date().getFullYear()}
        {'.'}
      </Typography>
    );
  }

  const useStyles = makeStyles(theme => ({
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      align: 'center',
      fontSize: '18px',
    },
    avatar: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(84),
      backgroundColor: theme.palette.secondary.main,
    },
    form: {
      width: '100%', // Fix IE 11 issue.
      marginTop: theme.spacing(1),
    },
    submit: {
      margin: theme.spacing(2, 0, 2),
    },
    textField: {
      width: 350 ,
      marginLeft:7,
      marginRight:-7,
     },
  
  }));

const Login: FC<{}> = () => {
  const classes = useStyles();

  const [User,setUser] = useState(String);
  const [Password,setPassword] = useState(String);

  const UserLoginhandlechange = (event : any) =>{
    setUser(event.target.value as string);
  }

  const Passwordhandlechange = (event : any) =>{
    setPassword(event.target.value as string);
  }
  return (
    <Grid container component="main" >
      <CssBaseline />
      <Grid item md={4} />

        <div className={classes.paper}>

          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            LOGIN TO SPORT CENTER
          </Typography>
          
          <Grid item xs={6}>
          <form className={classes.form} noValidate>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              autoComplete="username"
              autoFocus
              value={User || ""}
              onChange={UserLoginhandlechange}

            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              value={Password || ""}
              onChange={Passwordhandlechange}
            />
            
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label="Remember me"
            />
            <Button
              fullWidth
              variant="contained"
              color="primary"

              component={RouterLink}
              to="/StadiumUI"
            >
              Sign In
            </Button>
            
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="#" variant="body2" >
                  {"Don't have an account? Sign Up"}

                </Link>
              </Grid>
            </Grid>
            <Box mt={5}>
              <Copyright />
            </Box>
          </form>
          </Grid>
        </div>
      </Grid>
  );
};

export default Login;