import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Button from '@material-ui/core/Button'
import Link from '@material-ui/core/Link'
import Typography from '@material-ui/core/Typography'
import ShoppingCart from '@material-ui/icons/ShoppingCart'
import GitHubIcon from './icons/GitHub'
import { withStyles } from '@material-ui/core/styles'
import withWidth from '@material-ui/core/withWidth'

const styles = theme => ({
  appbar: {
    color: theme.palette.blackcard.text,
    backgroundColor: theme.palette.blackcard.background,
  },
  title: {
    flexGrow: 1,
  },
  icon: {
    margin: theme.spacing.unit,
    color: theme.palette.grey[50],
  },
});

function TopAppBar ({title, shortTitle, width, classes}) {
  return (
    <div>
      <AppBar position='static' className={classes.appbar} >
        <Toolbar>
          <Typography variant='h6' color='inherit'className={classes.title} >
            {width === "xs" ? shortTitle : title}
          </Typography>
          <Typography>
            <Link target="_blank" href="https://github.com/J4RV">
              <GitHubIcon className={classes.icon} />
            </Link>
          </Typography>
          <Typography> 
            <Link target="_blank" href="https://store.cardsagainsthumanity.com">
              <ShoppingCart className={classes.icon} />
            </Link>
          </Typography>
          <a href='user/logout'>
            <Button color='inherit'>Log out</Button>
          </a>
        </Toolbar>
      </AppBar>
    </div>
  )
}

export default withWidth()(withStyles(styles)(TopAppBar))
