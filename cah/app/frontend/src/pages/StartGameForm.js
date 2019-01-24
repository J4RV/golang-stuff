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
    textAlign: "left",
  },
  form: {
    marginTop: theme.spacing.unit * 2,
    display: "inline-block",
    width: "100%",
  },
})

class StartGameForm extends Component {
  render() {
    const { classes, gameID } = this.props
    return (
      <div className={classes.container}>
        <Typography variant="h6" gutterBottom>
          Game options
        </Typography>
        <form className={classes.form} action={startGameUrl} method="post">
          <input type="hidden" id="gameID" name="gameID" value={gameID} />
          <ExpansionsSelect />
          <FormControl required fullWidth margin="normal">
            <TextField
              id="handSize"
              name="handSize"
              label="Hand size"
              defaultValue={10}
              type="number"
            />
          </FormControl>
          <FormControl fullWidth margin="normal">
            <FormControlLabel
              control={<Checkbox id="randomFirstCzar" name="randomFirstCzar" color="primary" value={true} />}
              label="First Czar chosen randomly"
            />
          </FormControl>
        </form>
      </div>
    )
  }
}

export default withStyles(styles)(StartGameForm)
