import React, { Component } from "react"

import Checkbox from "@material-ui/core/Checkbox"
import ExpansionsSelect from "../components/ExpansionsSelect"
import FormControl from "@material-ui/core/FormControl"
import FormControlLabel from "@material-ui/core/FormControlLabel"
import { TextField } from "@material-ui/core"
import Typography from "@material-ui/core/Typography"
import { startGameUrl } from "../restUrls"
import { withStyles } from "@material-ui/core/styles"

const styles = theme => ({
  container: {
    padding: theme.spacing.unit * 2,
    maxWidth: 480,
    marginLeft: "auto",
    marginRight: "auto",
  },
  title: {
    textAlign: "center",
  },
  form: {
    marginTop: theme.spacing.unit * 2,
    display: "inline-block",
    textAlign: "right",
    width: "100%",
  },
})

class StartGameForm extends Component {
  render() {
    const { classes, gameID } = this.props
    return (
      <div className={classes.container}>
        <Typography variant="h6" gutterBottom className={classes.title}>
          Game options
        </Typography>
        <form className={classes.form} action={startGameUrl}>
          <input type="hidden" id="gameID" value={gameID} />
          <ExpansionsSelect />
          <FormControl fullWidth margin="normal">
            <TextField
              id="handSize"
              label="Hand size"
              value={10}
              type="number"
            />
          </FormControl>
          <FormControl fullWidth margin="normal">
            <FormControlLabel
              control={<Checkbox id="randomFirstCzar" color="primary" />}
              label="Random first Czar"
            />
          </FormControl>
        </form>
      </div>
    )
  }
}

export default withStyles(styles)(StartGameForm)
