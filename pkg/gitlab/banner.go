package gitlab

// Banner for clone all
const Banner = `
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN00KXXXXXK00NMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWx:oxxxxxxxockMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWdlXMMMMMMMNdxWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNKOxolc'cXMMMMMMMNl,cloxOKNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKko:''',;cl;cXMMMMMMMNl,lc;,'',:lkKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKd;..,lxOKNWMWxlXMMMMMMMNodWMWNX0xl;'';dKWMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMNOc.'cxKWMMMMMMMWxlXMMMMMMMNodWMMMMMMMWKkc,'cOWMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMW0c';xXWMMMMMMMMMMWxlXMMMMMMMNodWMMMMMMMMMMWXk:'lKWMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMWx''xNMMMMMMMMMMMMMWxlXMMMMMMMNodWMMMMMMMMMMMMMWk,,kWMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMWd.;0WMMMMMMMMMMMMMMWOxNMMMMMMMNkkWMMMMMMMMMMMMMMMK:'xWMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMWk.;KMMMMMMMMMMMMMMMMMWWMMMMMMMMMWWMMMMMMMMMMMMMMMMMK:'kWMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMK;,0MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0,:KMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMWd'xWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWx,xWMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMNccXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX:lNMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMXdkWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWkdXMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMNOo:lKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXo:okXMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMk.  '0MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX;  .dWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMk.  .xNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWO'   dWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMk.   .,;:cccccccccccccccccccccccccccccccccccccccccccccccc:,.   .dWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMk. ;dl:;;;,,,'''................................'''',,;;;:cld: .dWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMO. oWKdddxxxxxkkkkkkkkkkkkkkkkkkOOkkkkkkkkkkkkkkkxxxxxxxddd0Nx..dWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMO..dWXkoc;'..    ...............................     ..,coxKWx..xMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMO..xMMMMMWX0xo:'.                               .':lxOXNWMMMMk..xMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMO..dWMMMMMMMMMWNKOdc,.                     .,cdkKNWMMMMMMMMMWx..xMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0' 'xXMMMMMMMMMMMMMMWXOd:.             .;okKNMMMMMMMMMMMMMMNx' .xMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0'   ,xNMMMMMMMMMMMMMMMMWXo.         .lKWMMMMMMMMMMMMMMMMNk;   .kMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0'     ,kNMMMMMMMMMMMMMMMMWo.        lNMMMMMMMMMMMMMMMMWO:.    .kMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0'       ;kNMMMMMMMMMMMMMMMK;       ,0MMMMMMMMMMMMMMMW0:.      .kMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0'   .,.  .:OWMMMMMMMMMMMMMWx.     .dWMMMMMMMMMMMMMW0c.  .,.   .OMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMM0,   .kXx:. .:OWMMMMMMMMMMMMX:     :XMMMMMMMMMMMMWKl. .,dKO'   .OMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMK,   .dWMW0o,..c0WMMMMMMMMMMWx.   .oWMMMMMMMMMMWKl. 'lONMMk.   .OMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMK,  'c0MMMMMNkc..l0WMMMMMMMXd:.   .:dKWMMMMMMWKo'.:xXWMMMMKo,. .OMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMXxoONWMMMMMMMMWKd:;oXMMMMXx,.o0o;:kO:,oKWMMMNd;;oKWMMMMMMMMMN0ddKMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMNXWMMMMMMMMMMMMMN0xKMMXx,,dx0WMWWMXkxl,oKWMXkONMMMMMMMMMMMMMWXXMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMWWWMNxoKMMMMMMMMMMMMMMMMMXx''lkXMMMMMMMMMW0d:'l0WMMMMMMMMMMMMMMMMKdxXMWWMMMMMMMMMMMMMM
MMMMMMMMMMMMMMKkXM0llKMMMMMMMMMMMMMMMXx,.;kNMMMMMMMMMMMMMWKo,'lKWMMMMMMMMMMMMMMKolOWKONMMMMMMMMMMMMM
MMMMMMMMMMMMMMKoxWXkOWMMMMMMMMMMMMMNx;,l0WMMMMMMMMMMMMMMMMMMXx:,l0WMMMMMMMMMMMMM0xXXddNMMMMMMMMMMMMM
MMMMMMMMMMMMMMWOcdXWWMMMMMMMMMMMMNkc:dXWMMMMMMMMMMMMMMMMMMMMMMWOl:o0WMMMMMMMMMMMWNKolKMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMKolkNMMMMMMMMMMNOolONMMMMMMMMMMMMMMMMMMMMMMMMMMMWKxldKWMMMMMMMMWXdcdXMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMWOoldKWMMMMMW0xxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOxkXMMMMMNOoloKWMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMW0dllx0NMMWXNMMMMMMMMMMMMMMWXKKKKKXWMMMMMMMMMMMMMMWXXMWNOdookXWMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMN0dolox0NMMMMMMMMMMMMMMMK:.....;0MMMMMMMMMMMMMMWN0xdodkKWMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMWXOdllox0NWMMMMMMMMMM0' .,. '0MMMMMMMMMMWXOxdodx0NMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMWXOdolodkKNMMMMMMXddKWXddXMMMMMWNKkxdddkKNMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWN0xoooodk0KNWWWWMWWWNXKOkxdddxOXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWX0kddddddxxxxxxxxxxkOXNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNNXXXNNNWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
`
