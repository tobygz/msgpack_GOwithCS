using MessagePack;
using MessagePack.Formatters;
using System;
using SharedData;
using System.Collections.Generic;
using MessagePack.Resolvers;

namespace msgpacktest
{
    internal class Program
    {
        static void Main(string[] args)
        {
            /*
            StaticCompositeResolver.Instance.Register(
                new List<IMessagePackFormatter>() {
                            new TBuffer.TBufferFormatter()
                },
                new List<IFormatterResolver>() {
                            BuiltinResolver.Instance,
                            //GeneratedResolver.Instance
                });
            */

            var hbp = new HttpBasePacket();
            hbp.uuid = "A";
            var bin = new byte[] { 1, 1, 1 };
            hbp.rawData = new TBuffer(bin, bin.Length);
            /*
            hbp.mySDic = new SortedDictionary<int, int>();
            hbp.mySDic.Add(2, 2);
            hbp.mySDic.Add(1, 1);
            hbp.mySDic.Add(3, 3);
            */

            //hbp.myDic.Add(2, 2);
            //var bin = new byte[] { 1, 1, 1 };
            //hbp.rawData = new TBuffer(bin, 3);
            //hbp.uuid = "8";
            //hbp.timeStamp = 999;
            //hbp.messageType = 777;
            //var bin = new byte[] { 1, 1, 1 };
            //hbp.rawData = new TBuffer(bin, bin.Length);
            byte[] bytesb = MessagePackSerializer.Serialize(hbp);
            Console.WriteLine("bytes length: " + bytesb.Length);
            //145 145 192
            for (int i = 0; i < bytesb.Length; i++)
            {
                Console.Write(bytesb[i]);
                Console.Write(" ");
            }
            Console.WriteLine();
            //145 145 192

            var hbp_d = new HttpBasePacket();
            hbp_d = MessagePackSerializer.Deserialize<HttpBasePacket>(bytesb);
            Console.WriteLine(hbp_d.uuid);
            for (int i = 0; i < hbp_d.rawData.Buffer.Length; i++)
            {
                Console.Write(hbp_d.rawData.Buffer[i]);
                Console.Write(" ");
            }
            Console.WriteLine();
            // => 145 161 56
        }
    }
}
