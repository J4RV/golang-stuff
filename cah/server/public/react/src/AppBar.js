import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

function ButtonAppBar(props) {
  return (
    <div>
      <AppBar position="static" color="secondary" >
        <Toolbar>
          <Typography variant="h6" color="inherit" style={{ flexGrow: 1 }}>
            {props.title}
          </Typography>
          <a href="user/logout">
            <Button color="inherit">Logout</Button>
          </a>
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default ButtonAppBar;