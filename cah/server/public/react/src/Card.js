import React from 'react'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  card: {
    position: "relative",
    display: "inline-block",
    margin: "0 0 -10px 0",
    padding: theme.spacing.unit,
    borderRadius: 10,
    textAlign: "center",
    fontFamily: '"Open Sans", "Roboto", "Helvetica", "Arial", sans-serif',
    fontWeight: "800",
    verticalAlign: "top",
  },
  text: {
    width: "8rem",
    height: "10rem",
    fontSize: "0.8rem",
    whiteSpace: "pre-wrap",
  },
  black: {
    color: theme.palette.blackcard.text,
    background: theme.palette.blackcard.background,
  },
  white: {
    color: theme.palette.whitecard.text,
    background: theme.palette.whitecard.background,
  },
  expansion: {
    position: "absolute",
    fontSize: "0.6rem",
    bottom: theme.spacing.unit,
    right: theme.spacing.unit * 2,
    marginLeft: theme.spacing.unit,
    color: theme.palette.expansion,
  },
  inTable: {
    boxShadow: theme.shadows[1],
  },
  floating: {
    boxShadow: theme.shadows[10],
  },
  glowing: {
    boxShadow: theme.lights.glow,
  },
});

const Card = (props) => {
  const { text, isBlack, elevated, glowing, expansion, className, classes, style, ...rest } = props
  let shadowClass
  console.log(glowing)
  if (glowing) {
    shadowClass = classes.glowing
  } else {
    shadowClass = elevated ? classes.floating : classes.inTable
  }
  const colorClass = isBlack ? classes.black : classes.white
  return <div
    style={{ transform: `rotate(${Math.random() * 5 - 2.5}deg)`, ...style }}
    className={`${className ? className : null} ${classes.card} ${colorClass} ${shadowClass}`}
    {...rest}
  >
    <div className={classes.text}>{text}</div>
    <div className={classes.expansion}>{expansion}</div>
  </div>
}

export default withStyles(styles)(Card)
