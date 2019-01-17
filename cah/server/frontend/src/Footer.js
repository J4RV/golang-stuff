import React from 'react'
import Typography from '@material-ui/core/Typography'
import ShoppingCart from '@material-ui/icons/ShoppingCart'
import GitHubIcon from './icons/GitHub'
import Link from '@material-ui/core/Link'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  linkContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  icon: {
    margin: theme.spacing.unit,
  },
});

const Footer = ({classes}) => {
  return <div className={classes.linkContainer}>
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
  </div>
}

export default withStyles(styles)(Footer)