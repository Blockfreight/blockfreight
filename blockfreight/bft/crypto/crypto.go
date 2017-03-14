// File: ./blockfreight/bft/crypto/crypto.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright 2017 Blockfreight, Inc.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

// Package crypto provides useful functions to sign BF_TX.
package crypto

import (
    // =======================
    // Golang Standard library
    // =======================
    "crypto/ecdsa"      // Implements the Elliptic Curve Digital Signature Algorithm, as defined in FIPS 186-3.
    "crypto/elliptic"   // Implements several standard elliptic curves over prime fields.
    "crypto/md5"        // Implements the MD5 hash algorithm as defined in RFC 1321.
    "crypto/rand"       // Implements a cryptographically secure pseudorandom number generator.
    "fmt"               // Implements formatted I/O with functions analogous to C's printf and scanf.
    "hash"              // Provides interfaces for hash functions.
    "io"                // Provides basic interfaces to I/O primitives.
    "math/big"          // Implements arbitrary-precision arithmetic (big numbers).
    "os"                // Provides a platform-independent interface to operating system functionality.
    "strconv"           // Implements conversions to and from string representations of basic data types.

    // ======================
    // Blockfreight™ packages
    // ======================
    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx" // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

//  @todo: OP_2 <pubkey1> <pubkey2> <pubkey3> <pubkey4> <pubkey5> OP_3 OP_CHECKMULTISIGVERIFY <pubkey3> OP_CHECKSIG

// Function Sign_BF_TX has the whole process of signing each BF_TX.
func Sign_BF_TX(bft_tx bf_tx.BF_TX) bf_tx.BF_TX {

    content := bf_tx.BF_TXContent(bft_tx)

    pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

    privatekey := new(ecdsa.PrivateKey)
    privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    pubkey := privatekey.PublicKey

    // Sign ecdsa style
    var h hash.Hash
    h = md5.New()
    r := big.NewInt(0)
    s := big.NewInt(0)

    io.WriteString(h, content)
    signhash := h.Sum(nil)

    r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
    if serr != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    signature := r.Bytes()
    signature = append(signature, s.Bytes()...)
    
    sign := ""
    for i, _ := range signature {
        sign += strconv.Itoa(int(signature[i]))
    }
    
    // Verification
    verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
    
    //Set Private Key and Sign to BF_TX
    bft_tx.PrivateKey = *privatekey
    bft_tx.Signhash = signhash
    bft_tx.Signature = sign
    bft_tx.Verified = verifystatus
    
    return bft_tx
}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================
