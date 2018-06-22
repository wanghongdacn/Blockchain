@echo off

SET "target=api"
echo Generating README.md for %target%
cd ..\%target%\
godocdown \github.com\NlaakStudios\Blockchain\%target% > README.md
cd ..

SET "target=cli"
echo Generating README.md for %target%
cd ..\%target%\
godocdown \github.com\NlaakStudios\Blockchain\%target% > README.md
cd ..

SET "target=config"
echo Generating README.md for %target%
cd .\%target%\
godocdown \github.com\NlaakStudios\Blockchain\%target% > README.md
cd ..

SET "target=core"
echo Generating README.md for %target%
cd .\%target%\
godocdown \github.com\NlaakStudios\Blockchain\%target% > README.md
cd ..

SET "target=utils
echo Generating README.md for %target%
cd .\%target%\
godocdown \github.com\NlaakStudios\Blockchain\%target% > README.md
cd ..

echo Done.