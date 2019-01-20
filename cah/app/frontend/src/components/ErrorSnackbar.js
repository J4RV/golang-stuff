import React, { Component } from 'react'
import ErrorIcon from '@material-ui/icons/Error'
import Snackbar from '@material-ui/core/Snackbar'
import SnackbarContent from '@material-ui/core/SnackbarContent'
import CloseIcon from '@material-ui/icons/Close'
import IconButton from "@material-ui/core/IconButton"
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  error: {
    color: theme.palette.getContrastText(theme.palette.error.dark),
    background: theme.palette.error.dark,
    display: 'flex',
    alignItems: 'center',
  },
  icon: {
    marginRight: theme.spacing.unit,
  },
  message: {
    display: 'flex',
    alignItems: 'center',
  },
});

class ErrorSnackbar extends Component {
  render() {
    const classes = this.props.classes
    return <Snackbar
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      open={this.props.msg != null && this.props.msg !== ""}
    >
      <SnackbarContent
        className={classes.error}
        aria-describedby="message-id"
        message={
          <span id="message-id" className={classes.message}>
            <ErrorIcon className={classes.icon} />
            {this.props.msg}
          </span>
        }
        action={[
        <IconButton
          key="close"
          aria-label="Close"
          color="inherit"
          onClick={this.props.onClose}
        >
          <CloseIcon />
        </IconButton>
      ]}
      />
    </Snackbar>
  }
}
export default withStyles(styles)(ErrorSnackbar)