import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button'
import GitHubIcon from './icons/GitHub'
import React from 'react'
import ShoppingCart from '@material-ui/icons/ShoppingCart'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { logoutUrl } from './restUrls'
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

function TopAppBar({ title, shortTitle, width, classes }) {
  return (
    <div>
      <AppBar position='static' className={classes.appbar} >
        <Toolbar>
          <Typography variant='h6' color='inherit' className={classes.title} >
            {width === "xs" ? shortTitle : title}
          </Typography>
          <Typography>
            <a target="blank" href="https://github.com/J4RV">
              <GitHubIcon className={classes.icon} />
            </a>
          </Typography>
          <Typography>
            <a target="blank" href="https://store.cardsagainsthumanity.com">
              <ShoppingCart className={classes.icon} />
            </a>
          </Typography>
          <a href={logoutUrl}>
            <Button color='inherit'>Log out</Button>
          </a>
        </Toolbar>
      </AppBar>
    </div>
  )
}

export default withWidth()(withStyles(styles)(TopAppBar))
