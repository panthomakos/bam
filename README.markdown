# BAM!

You have personal projects, you have work projects. You have a lot of directories and you need to navigate to them quickly... BAM! and you're there. BAM! is a simple and speedy utility to quickly navigate to the root of project directories using a fuzzy prefix syntax.

## Installation

1. Download a recent release of BAM! from the [GitHub releases page](https://github.com/panthomakos/bam/releases).
2. Unpack it: `tar -xvf <file>.tar.gz`.
3. Move the `bam` binary to `/usr/local/bin`: `mv bam /usr/local/bin/`.
4. Configure an alias with your default configuration options:

        # .zshrc or .bashrc
        function bam() {
          local dest=$(${HOME}/bin/bam -root $HOME/Projects $1)
          cd $dest
        }

## Example Usage

    // Directories:

    $HOME/Projects/science/physics/black-holes
    $HOME/Projects/science/physics/nebulae
    $HOME/Projects/ballet
    $HOME/Projects/bingo

    // Quickly navigate to the black-holes directory by indicating
    // that you are searching for a directory whose name starts with
    // "b" that is nested somewhere beneath a directory whose name
    // starts with "s".
    bam s/b

    // Quickly navigate to the "ballet" directory because it is sorted
    // before the "bingo" directory.
    bam b

    // Quickly navigate to the "bingo" directory.
    bam bin

## TODO

You name it, right now we're alpha but it works for me :)
