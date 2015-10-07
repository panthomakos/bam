# BAM!

A simple and speedy utility to quickly navigate to the root of project
directories using a prefix syntax.

## Example:

    // Directories:

    $HOME/Projects/science/physics/black-holes
    $HOME/Projects/ballet
    $HOME/Projects/bingo

    // Quickly navigate to the first directory.
    bam s/b

    // Quickly navigate to the second directory.
    bam b

    // Quickly navigate to the third directory.
    bam bin

## Configured in ZSHRC

    function bam() {
      local dest=$(${HOME}/bin/bam $1)
      cd $dest
    }

## TODO

You name it, right now it's pre-alpha but it works for local development.
