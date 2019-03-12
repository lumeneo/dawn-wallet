# Dawn Wallet

The universal GUI wallet for Lumeneo.

## Installation

[Windows](#windows) - [Mac](#mac) - [Linux](#linux)

### Windows

1. Go [here](https://github.com/lumeneo-network/dawn-wallet/releases) and download the latest release called **dawn-wallet-x.xx-Windows.zip**
2. Unzip the folder and launch **dawn-wallet.exe**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *Walletd.exe* or *Lumeneod.exe* are not running before you start *dawn-wallet*
* When you open a wallet in Dawn, you will see a black empty *walletd* window. You must keep it open. It will close automatically when you close your wallet. If it does not close automatically a few seconds after you close your wallet, you can close it manually.

### Mac

1. Go [here](https://github.com/lumeneo-network/dawn-wallet/releases) and download the latest release called **dawn-wallet-x.xx-Mac.zip**.
2. Unzip it and move the folder wherever you want or drag the application **dawn-wallet** into /Applications or any other folder.
3. Launch the application. (If your mac complains that the app comes from an unindentified developer and does not want to open it, just right-click (or ctrl-click) on the app, and choose open > open)

Important notes:

* The wallets you create or generate will be saved to your home folder. You can keep them there or move them wherever you want.
* Make sure *Walletd* or *Lumeneod* are not running before you start *dawn-wallet*.
* If you encounter crashes, open the activity monitor (in your app > utilities), and force quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on).
* The log files will be saved in ~/Library/Application Support/dawn-wallet/.

### Linux

1. Go [here](https://github.com/lumeneo-network/dawn-wallet/releases) and download the latest release called **dawn-wallet-x.xx-Linux.tar.gz**
2. extract it
`$ tar xvzf dawn-wallet-x.xx-Linux.tar.gz`
3. run **dawn-wallet.sh**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *Walletd* or *Lumeneod* are not running before you start *dawn-wallet*
* If you want the *copy address to clipboard* button to work, install *xclip* or *xsel* (on Debian/Ubuntu: `$ sudo apt install xclip`).
* If you encounter crashes, open an activity monitor (e.g. `$ htop`), and quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on)

## Upgrade

Just download the new release and follow the same steps as [Installation](#installation). Just make sure you don't delete your wallet files in the old folder (.wallet files) and you copy them or move them to a new folder.

## Donations

TRTLv3jzutiQwqHL3qFwsu5EVLWesxZr1AFQ4AuMR3SD56n3rkHDkwj79eKwvaiU1nYQWGydKoXM6fXyiiGKsPDnVCNXzNdusxx

## Build

(for developers only)

1. Install this binding: https://github.com/therecipe/qt
2. Run `qtdeploy build desktop`
3. The app folder is in deploy/*your os*/
3. Include the latest walletd build in:
    * Windows, Linux: in the app folder
    * Mac: in dawn-wallet.app/Contents/
